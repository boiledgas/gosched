package main

import (
	"gosched.git"
	"log"
	"runtime"
	"sync"
	"time"
)

/*
1. Не обрабатывается период на конце
2. Перенести период в интервал
3. Реализовать количество возможное повторений
*/
func main() {
	runtime.GOMAXPROCS(8)
	s := gosched.Schedule{
		Conditions: make(map[gosched.ConditionId]gosched.Condition),
		Relations:  make(map[gosched.RelationKey]gosched.Relation),
	}
	//15 10-16:^15:14
	//15 15-21:^30:13
	root := gosched.ROOT_ID
	d1Id, _ := s.Add(root, gosched.Condition{Part: gosched.CONDITION_PART_DAY, Type: gosched.CONDITION_EXACT, Value: 15})
	h1Id, _ := s.Add(d1Id, gosched.Condition{Part: gosched.CONDITION_PART_HOUR, Type: gosched.CONDITION_BOUND, From: 10, To: 16})
	h2Id, _ := s.Add(d1Id, gosched.Condition{Part: gosched.CONDITION_PART_HOUR, Type: gosched.CONDITION_BOUND, From: 11, To: 15})
	/*m1Id, _ :=*/ s.Add(h1Id, gosched.Condition{Part: gosched.CONDITION_PART_MIN, Type: gosched.CONDITION_REPEAT, Value: 15})
	/*m2Id, _ :=*/ s.Add(h2Id, gosched.Condition{Part: gosched.CONDITION_PART_MIN, Type: gosched.CONDITION_REPEAT, Value: 30})
	//s.Add(m1Id, gosched.Condition{Part: gosched.CONDITION_PART_SEC, Type: gosched.CONDITION_EXACT, Value: 14})
	//s.Add(m2Id, gosched.Condition{Part: gosched.CONDITION_PART_SEC, Type: gosched.CONDITION_EXACT, Value: 13})

	time := time.Date(2008, 1, 1, 0, 0, 13, 0, time.UTC)
	ok := true
	state := gosched.ScheduleState{
		Cache: make(map[gosched.ConditionId]gosched.ConditionState),
	}
	i := 0
	for time, ok = s.Handle(time, &state); ok ; time, ok = s.Handle(time, &state) {
		log.Printf("time: %v", time)
		log.Println("--------------------")
		i++
	}
}

func performance(s *gosched.Schedule) {
	logScheduler(s)
	wg := &sync.WaitGroup{}
	duration := time.Now().UnixNano()
	for i := 0; i < 50; i++ {
		wg.Add(1)
		func() {
			Add(s, 0, 0, wg)
			wg.Done()
		}()
	}
	wg.Wait()
	duration = time.Now().UnixNano() - duration
	log.Printf("relations: %v conditions: %v proxies: %v time(%v)", s.CountRelations(), s.Count(), s.CountProxies(), duration)
	log.Printf("%v", s.String())
}

func logScheduler(s *gosched.Schedule) {
	log.Printf("relations: %v conditions: %v proxies: %v", s.CountRelations(), s.Count(), s.CountProxies())
	time.AfterFunc(time.Second*1, func() {
		logScheduler(s)
	})
}

func Add(s *gosched.Schedule, parentId gosched.ConditionId, level uint16, wg *sync.WaitGroup) (err error) {
	if level == 3 {
		return
	}
	for i := 0; i < 5; i++ {
		var pid gosched.ConditionId
		if pid, err = s.Add(parentId, gosched.Condition{Value: byte(i)}); err != nil {
			log.Printf("add: %v", err)
			break
		}
		Add(s, pid, level+1, wg)
	}
	return
}
