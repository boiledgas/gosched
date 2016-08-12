package gosched

import "testing"

func BenchmarkSchedule(b *testing.B) {
	s := Schedule{
		Conditions: make(map[ConditionId]Condition),
		Relations:  make(map[RelationKey]Relation),
	}
	for i := 0; i < b.N; i++ {
		s.Add(0, Condition{Value:byte(i)})
	}
}
