package gosched

import (
	"fmt"
	"log"
	"testing"
	"time"
	"unsafe"
)

func _Test_Sizeof(t *testing.T) {
	r := restriction_container{}
	b1 := make([][]byte, 1)
	b2 := make([]restriction, 1)
	b3 := restriction{}
	log.Println(fmt.Sprintf("bytes: %v %v %v %v", unsafe.Sizeof(r), unsafe.Sizeof(b1), unsafe.Sizeof(b2), unsafe.Sizeof(b3)))
}

type Test struct {
}

func _Test_builder(t *testing.T) {
	r := NewRestriction()
	b := NewBuilder(r)
	b.YearBetweenAll(16, 18, func(b RestrictionBuilder) {
	})

	b.YearToAll(16, func(b RestrictionBuilder) {
		b.MonthFromAll(2, func(b RestrictionBuilder) {
			b.Day(2).Hour(10)
			b.Day(10).Hour(14).Minute(30)
		})
	})

	b.MonthAll(2, func(b RestrictionBuilder) {
		b.Day(25).HourEvery(7)
	})

	log.Println(r)
}

func Test_Schedule_Basic1(t *testing.T) {
	r := NewRestriction()

	b := NewBuilder(r)
	b.MonthBetween(4, 10).MonthEvery(2).DayEveryAll(2, func(b RestrictionBuilder) {
		b.HourBetween(12, 14).MinuteEvery(30)
		b.HourFrom(16).MinuteFrom(0).MinuteEvery(30).Second(40)
	})

	start := time.Date(2016, 7, 30, 16, 40, 0, 0, time.UTC)
	result := start
	var ok bool
	for i := 0; i < 21; i++ {
		result, ok = r.Handle(result)
		log.Printf("result %v %v", result, ok)
		if !ok {
			break
		}
	}
	//	result, _ := r.Handle(start)
	//	if result.Unix() != time.Date(2016, 9, 10, 10, 0, 0, 0, time.UTC).Unix() {
	//		t.Fail()
	//	}
	//	result, _ = r.Handle(result)
	//	if result.Unix() != time.Date(2016, 9, 10, 10, 10, 0, 0, time.UTC).Unix() {
	//		t.Fail()
	//	}
	//	result, _ = r.Handle(result)
	//	if result.Unix() != time.Date(2016, 9, 10, 10, 20, 0, 0, time.UTC).Unix() {
	//		t.Fail()
	//	}
}

func _Test_Schedule(t *testing.T) {
	r := NewRestriction()
	b := NewBuilder(r)

	b.Month(6).DayEveryAll(1, func(b RestrictionBuilder) {
		b.Hour(10)
		b.HourBetweenAll(12, 13, func(b RestrictionBuilder) {
			b.Minute(10)                    // точно в 10 минут (самый высокий приоритет)
			b.MinuteFrom(30).MinuteEvery(5) // с 30 минут - каждые 5 минут
			b.MinuteFrom(55).SecondEvery(5) // с 55 минут должно страбатывать это условие
		})
	})
}
