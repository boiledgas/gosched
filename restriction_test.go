package gosched

import (
	"fmt"
	"log"
	"testing"
	"unsafe"
)

func Test_Sizeof(t *testing.T) {
	r := restriction_container{}
	b1 := make([][]byte, 1)
	b2 := make([]restriction, 1)
	b3 := restriction{}
	log.Println(fmt.Sprintf("bytes: %v %v %v %v", unsafe.Sizeof(r), unsafe.Sizeof(b1), unsafe.Sizeof(b2), unsafe.Sizeof(b3)))
}

type Test struct {
}

type BuilderFunc func(b RestrictionBuilder)
type year_restriction_builder interface {
	BYear(byte, byte) RestrictionBuilder // between year
	FYear(byte) RestrictionBuilder       // from year
	TYear(byte) RestrictionBuilder       // to year
	EYear(byte) RestrictionBuilder       // every year
	Year(byte) RestrictionBuilder        // this year

	ABYear(byte, byte, BuilderFunc) // all between year
	AFYear(byte, BuilderFunc)       // all from year
	ATYear(byte, BuilderFunc)       // all to year
	AEYear(byte, BuilderFunc)       // all every year
	AYear(byte, BuilderFunc)        // all this year
}
type month_restriction_builder interface {
	MonthBetween(byte, byte) RestrictionBuilder
	MonthFrom(byte) RestrictionBuilder
	MonthTo(byte) RestrictionBuilder
	EveryMonth(byte) RestrictionBuilder
	Month(byte) RestrictionBuilder

	MonthBetweenAll(byte, byte, BuilderFunc)
	MonthFromAll(byte, BuilderFunc)
	MonthToAll(byte, BuilderFunc)
	EveryMonthAll(byte, BuilderFunc)
	MonthAll(byte, BuilderFunc)
}
type day_restriction_builder interface {
	DayBetween(byte, byte) RestrictionBuilder
	DayFrom(byte) RestrictionBuilder
	DayTo(byte) RestrictionBuilder
	EveryDay(byte) RestrictionBuilder
	Day(byte) RestrictionBuilder

	DayBetweenAll(byte, byte, BuilderFunc)
	DayFromAll(byte, BuilderFunc)
	DayToAll(byte, BuilderFunc)
	EveryDayAll(byte, BuilderFunc)
	DayAll(byte, BuilderFunc)
}
type hour_restriction_builder interface {
	HourBetween(byte, byte) RestrictionBuilder
	HourFrom(byte) RestrictionBuilder
	HourTo(byte) RestrictionBuilder
	EveryHour(byte) RestrictionBuilder
	Hour(byte) RestrictionBuilder

	HourBetweenAll(byte, byte, BuilderFunc)
	HourFromAll(byte, BuilderFunc)
	HourToAll(byte, BuilderFunc)
	EveryHourAll(byte, BuilderFunc)
	HourAll(byte, BuilderFunc)
}
type minute_restriction_builder interface {
	MinuteBetween(byte, byte) RestrictionBuilder
	MinuteFrom(byte) RestrictionBuilder
	MinuteTo(byte) RestrictionBuilder
	EveryMinute(byte) RestrictionBuilder
	Minute(byte) RestrictionBuilder

	MinuteBetweenAll(byte, byte, BuilderFunc)
	MinuteFromAll(byte, BuilderFunc)
	MinuteToAll(byte, BuilderFunc)
	EveryMinuteAll(byte, BuilderFunc)
	MinuteAll(byte, BuilderFunc)
}
type second_restriction_builder interface {
	SecondBetween(byte, byte) RestrictionBuilder
	SecondFrom(byte) RestrictionBuilder
	SecondTo(byte) RestrictionBuilder
	EverySecond(byte) RestrictionBuilder
	Second(byte) RestrictionBuilder

	SecondBetweenAll(byte, byte, BuilderFunc)
	SecondFromAll(byte, BuilderFunc)
	SecondToAll(byte, BuilderFunc)
	EverySecondAll(byte, BuilderFunc)
	SecondAll(byte, BuilderFunc)
}

type RestrictionBuilder interface {
	year_restriction_builder
	month_restriction_builder
	day_restriction_builder
	hour_restriction_builder
	minute_restriction_builder
	second_restriction_builder
}

func Test_builder(t *testing.T) {
	b := RestrictionBuilder(nil)
	b.YearBetweenAll(16, 18, func(b RestrictionBuilder) {
	})

	b.YearFromAll(16, func(b RestrictionBuilder) {
		b.MonthFromAll(2, func(b RestrictionBuilder) {
			b.EveryMonth(2).Hour(10)
			b.EveryHour(10).Hour(14).Minute(30)
		})
	})

	b.MonthAll(2, func(b RestrictionBuilder) {
		b.Day(25).EveryHour(7)
	})

}
