package gosched

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"time"
)

const ROOT_ID ConditionId = 0
const RELATION_SIZE = 5

type ForEachFunc func(parentId ConditionId, id ConditionId, c Condition)

type RelationType byte

const (
	RELATION_NULL      RelationType = 0x00 // relation is null
	RELATION_PROXY     RelationType = 0x01 // relation is pointer to another relation
	RELATION_CONDITION RelationType = 0x02 // relation is pointer to condition
)

type RelationKey struct {
	Type RelationType // relation type
	Id   ConditionId  // relation identity
}

var EMPTY_RELATION_KEY RelationKey = RelationKey{
	Type: RELATION_NULL,
	Id:   0,
}

type Relation struct {
	Children [RELATION_SIZE]RelationKey
}

type HandleResult struct {
	Id     ConditionId
	Time   time.Time
	Cached bool
}

type Schedule struct {
	Conditions map[ConditionId]Condition
	Relations  map[RelationKey]Relation
	States     map[ConditionId]ConditionState
	proxyCount ConditionId
}

func (s *Schedule) Add(parentId ConditionId, c Condition) (id ConditionId, err error) {
	id = s.getConditionId()
	if id == 0 {
		err = errors.New("stack overflow")
		return
	}
	s.Conditions[id] = c
	if parentId == 0 {
		s.setRootRelation(id)
	} else {
		s.setConditionRelation(parentId, id)
	}
	return
}

func (s *Schedule) getConditionId() ConditionId {
	return ConditionId(len(s.Conditions) + 1)
}

func (s *Schedule) Remove(conditionId ConditionId) (err error) {
	var ok bool
	if _, ok = s.Conditions[conditionId]; !ok {
		return errors.New("not exists")
	}
	relationKey := RelationKey{
		Id:   conditionId,
		Type: RELATION_CONDITION,
	}
	err = s.removeRelation(relationKey)
	return
}

func (s *Schedule) ForEach(eachFunc ForEachFunc) (err error) {
	rootKey := RelationKey{Id: 0, Type: RELATION_PROXY}
	return s.foreach(0, rootKey, eachFunc)
}

func (s *Schedule) Count() int {
	return len(s.Conditions)
}

func (s *Schedule) CountRelations() int {
	return len(s.Relations)
}

func (s *Schedule) CountProxies() int {
	return int(s.proxyCount)
}

func (s *Schedule) Handle(currentTime time.Time, state *ScheduleState, metric *Metric) (result time.Time, ok bool) {
	rootKey := RelationKey{Id: 0, Type: RELATION_PROXY}
	var handleResult HandleResult
	//log.Printf("process: %v", currentTime)
	handleResult, ok = s.handleRelation(rootKey, currentTime, state, metric)
	result = handleResult.Time
	return
}

func (s *Schedule) handleRelation(key RelationKey, computeTime time.Time, state *ScheduleState, metric *Metric) (result HandleResult, ok bool) {
	var relation Relation
	var relationExists bool
	if relation, relationExists = s.Relations[key]; !relationExists {
		result.Time = computeTime
		ok = true
		return
	}

	ok = false
	empty := 0
	var handleOk bool
	var handleResult HandleResult
	for _, relationKey := range relation.Children {
		switch relationKey.Type {
		case RELATION_CONDITION:
			handleResult, handleOk = s.handleCondition(relationKey, computeTime, state, metric)
			//log.Printf("- condition %v %v %v", relationKey, computeTime, handleResult)
		case RELATION_PROXY:
			handleResult, handleOk = s.handleRelation(relationKey, computeTime, state, metric)
			//log.Printf("- proxy %v %v %v", relationKey, computeTime, handleResult)
		case RELATION_NULL:
			empty++
			continue
		}

		if handleOk {
			minFound := handleResult.Time.Unix() < result.Time.Unix()
			switch {
			case !minFound && ok && !handleResult.Cached:
				metric.CacheWrite()
				state.SetCache(handleResult.Id, handleResult.Time)
			case minFound && ok && !result.Cached:
				metric.CacheWrite()
				state.SetCache(result.Id, result.Time)
			}
			if minFound || !ok {
				//log.Printf("- min %v %v %v", key, relationKey, handleResult)
				result = handleResult
				ok = true
			}
		}
	}
	ok = ok || empty == len(relation.Children)
	if ok && result.Cached { // and cached
		metric.CacheDelete()
		state.Remove(result.Id)
	}
	//log.Printf("handle: %v %v %v", relationKey, result, ok)
	return
}

func (s *Schedule) handleCondition(key RelationKey, lastTime time.Time, state *ScheduleState, metric *Metric) (result HandleResult, ok bool) {
	condition := s.Conditions[key.Id]
	if time, stateOk := state.Get(key.Id); stateOk {
		metric.CacheRead()
		result.Time = time
		result.Id = key.Id
		result.Cached = true
		ok = true
		return
	}

	bounded, ok := condition.Bound(lastTime)
	if !ok {
		return
	}
	//	log.Printf("bound: %v %v => %v", key, lastTime, bounded)
	if result, ok = s.handleRelation(key, bounded, state, metric); ok {
		ok = result.Time.Unix() != lastTime.Unix()
	}
	if !ok {
		if repeatTime, repeated := condition.Repeat(bounded); repeated {
			result, ok = s.handleRelation(key, repeatTime, state, metric)
		}
	}

	result.Id = key.Id
	ok = condition.Check(result.Time)
	//log.Printf("handle: %v %v => %v %v", key, lastTime, result, ok)
	return
}

func (s *Schedule) foreach(parentId ConditionId, key RelationKey, eachFunc ForEachFunc) (err error) {
	switch key.Type {
	case RELATION_CONDITION:
		if condition, ok := s.Conditions[key.Id]; !ok {
			err = errors.New("not found")
		} else {
			eachFunc(parentId, key.Id, condition)
			parentId = key.Id
		}
	case RELATION_NULL:
		return
	}
	if relation, ok := s.Relations[key]; ok {
		for _, childKey := range relation.Children {
			if err = s.foreach(parentId, childKey, eachFunc); err != nil {
				return
			}
		}
	}
	return
}

func (s *Schedule) setRootRelation(id ConditionId) (err error) {
	var ok bool
	rootKey := RelationKey{Type: RELATION_PROXY, Id: 0}
	if _, ok = s.Relations[rootKey]; !ok {
		s.Relations[rootKey] = Relation{}
	}
	err = s.setRelation(rootKey, id)
	return
}

func (s *Schedule) setConditionRelation(parentId ConditionId, id ConditionId) (err error) {
	parentKey := RelationKey{Id: parentId, Type: RELATION_CONDITION}
	var ok bool
	if _, ok = s.Relations[parentKey]; !ok {
		s.Relations[parentKey] = Relation{}
	}
	err = s.setRelation(parentKey, id)
	return
}

func (s *Schedule) setRelation(key RelationKey, id ConditionId) (err error) {
	var index int
	parentRelation := s.Relations[key]
	for index = 0; index < len(parentRelation.Children); index++ {
		if parentRelation.Children[index].Type == RELATION_NULL {
			break
		}
	}
	conditionKey := RelationKey{Id: id, Type: RELATION_CONDITION}
	if index < len(parentRelation.Children) {
		// add relation to end
		parentRelation.Children[index] = conditionKey
	} else {
		// create proxy to relation.Children
		s.proxyCount++
		if s.proxyCount == 0 {
			return errors.New("stack overflow")
		}
		proxyKey := RelationKey{Type: RELATION_PROXY, Id: s.proxyCount}
		s.Relations[proxyKey] = parentRelation

		parentRelation.Children[0] = proxyKey
		parentRelation.Children[1] = conditionKey
		for i := 2; i < len(parentRelation.Children); i++ {
			parentRelation.Children[i] = EMPTY_RELATION_KEY
		}
	}
	s.Relations[key] = parentRelation
	return
}

func (s *Schedule) removeRelation(key RelationKey) (err error) {
	var relation Relation
	var ok bool
	if relation, ok = s.Relations[key]; !ok {
		err = errors.New("not exists")
		return
	}
	for _, childKey := range relation.Children {
		switch childKey.Type {
		case RELATION_PROXY:
			if err := s.removeRelation(childKey); err != nil {
				log.Printf("remove: %v", err)
			}
		case RELATION_CONDITION:
			delete(s.Conditions, childKey.Id)
		case RELATION_NULL:
		}
	}
	delete(s.Relations, key)
	return
}

func (s *Schedule) String() (str string) {
	var buf bytes.Buffer
	var value byte
	buf.WriteString("Relations:\n")
	for key, relation := range s.Relations {
		var childBuf bytes.Buffer
		for _, child := range relation.Children {

			if child.Type == RELATION_CONDITION {
				value = s.Conditions[child.Id].Value
			} else {
				value = 0
			}
			childBuf.WriteString(fmt.Sprintf(" (%v,%v,%v)", child.Id, child.Type.String(), value))
		}
		buf.WriteString(fmt.Sprintf("(%v,%v) :%v\n", key.Id, key.Type.String(), childBuf.String()))
	}
	buf.WriteString("Conditions:\n")
	for key, condition := range s.Conditions {
		buf.WriteString(fmt.Sprintf("%v - %v\n", key, condition))
	}
	str = buf.String()
	return
}

func (t RelationType) String() string {
	switch t {
	case RELATION_NULL:
		return "null"
	case RELATION_PROXY:
		return "proxy"
	case RELATION_CONDITION:
		return "condition"
	}
	return ""
}
