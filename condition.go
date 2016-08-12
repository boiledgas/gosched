package gosched

import (
	"time"
	"log"
)

type ConditionId uint32

type ConditionType byte

const (
	CONDITION_UNKNOWN ConditionType = 0x00
	CONDITION_BOUND   ConditionType = 0x01
	CONDITION_REPEAT  ConditionType = 0x02
	CONDITION_EXACT   ConditionType = 0x03
)

type Condition struct {
	Part  ConditionPart // part of date
	Type  ConditionType // null, bound, period, exact
	From  byte          // from bound value
	To    byte          // to bound value
	Value byte          // value
}

type ScheduleState struct {
	Cache map[ConditionId]ConditionState // condition states
}

func (s *ScheduleState) SetCache(id ConditionId, time time.Time) {
	log.Printf("- cached %v %v", id, time)
	var ok bool
	var state ConditionState
	if state, ok = s.Cache[id]; !ok {
		state = ConditionState{}
	}
	state.Value = time
	s.Cache[id] = state
}

func (s *ScheduleState) Remove(id ConditionId) {
	log.Printf("- cache clear %v", id)
	if _, ok := s.Cache[id]; ok {
		delete(s.Cache, id)
	}
}

func (s *ScheduleState) Get(id ConditionId) (time time.Time, ok bool) {
	log.Printf("- cache value: %v %v", id, time)
	var state ConditionState
	if state, ok = s.Cache[id]; ok {
		time = state.Value
	}
	return
}

type ConditionState struct {
	Value time.Time // cached value
}

func (c *Condition) Bound(time time.Time) (result time.Time, ok bool) {
	switch c.Type {
	case CONDITION_BOUND:
		if c.From > c.To {
			return
		}
		from_time := c.Part.Set(time, c.From)
		if time.Unix() < from_time.Unix() {
			result = from_time
		} else {
			result = time
		}
	case CONDITION_EXACT:
		if c.Part.Get(time) != c.Value {
			result = c.Part.Set(time, c.Value)
		} else {
			result = time
		}
	default:
		result = time
	}
	ok = result.Unix() >= time.Unix()
	return
}

func (c *Condition) Check(time time.Time) (ok bool) {
	unixTime := time.Unix()
	switch c.Type {
	case CONDITION_BOUND:
		if c.From < c.To {
			ok = unixTime >= c.Part.Set(time, c.From).Unix()
		}
		if c.To > 0 {
			ok = ok && unixTime <= c.Part.Set(time, c.To).Unix()
		}
	case CONDITION_EXACT:
		part := c.Part.Get(time)
		ok = part == c.Value
	default:
		ok = true
	}
	return
}

func (c *Condition) Repeat(time time.Time) (result time.Time, ok bool) {
	ok = c.Type == CONDITION_REPEAT
	if ok {
		result = c.Part.Set(time, c.Part.Get(time)+c.Value)
	}
	return
}

func (c *Condition) Fixup(time time.Time, state ConditionState) {
	// if time > state. state.Time()
}
