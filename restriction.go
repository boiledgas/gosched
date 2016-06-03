package gosched

import (
	"container/list"
	"fmt"
	"time"
)

type TimerType byte

const (
	TIMER_NULL     TimerType = 0x00
	TIMER_PERIOD   TimerType = 0x01
	TIMER_TIME     TimerType = 0x02
	TIMER_INTERVAL TimerType = 0x03
)

type DatePart byte

const (
	DATEPART_YEAR   DatePart = 0x01
	DATEPART_MONTH  DatePart = 0x02
	DATEPART_DAY    DatePart = 0x03
	DATEPART_HOUR   DatePart = 0x04
	DATEPART_MINUTE DatePart = 0x05
	DATEPART_SECOND DatePart = 0x06
)

// size 48
type restriction_container struct {
	rootId       byte
	relations    [][]byte      // heap
	restrictions []restriction // heap
}

const (
	RESTRICTION_NULL     byte = 0x00
	RESTRICTION_INTERVAL byte = 0x01
	RESTRICTION_TIME     byte = 0x02
	RESTRICTION_PERIOD   byte = 0x03
)

type restriction struct {
	r_type TimerType // null, interval, period, time
	part   DatePart  // part of date
	value  byte      // period, exact, operator
	from   byte      // interval from
	to     byte      // interval to
}

func (c *restriction_container) Handle(restrictionId byte, t time.Time) (res bool, time time.Time) {
	res = false
	time = t

	r_ids := c.relations[restrictionId]
	for _, r_id := range r_ids {
		if temp, temp_time := c.handleRestriction(r_id, t); temp && (time.Unix() < temp_time.Unix() || !res) {
			time = temp_time
			res = true
		}
	}

	return
}

func (c *restriction_container) handleRestriction(id byte, t time.Time) (res bool, time time.Time) {
	r := c.restrictions[id]
	switch r.r_type {
	case TIMER_INTERVAL:
		return c.handleInterval(id, r, t)
	case TIMER_PERIOD:
		return c.handlePeriod(id, r, t)
	case TIMER_TIME:
		return c.handleExact(id, r, t)
	default:
		panic("unknown r_type")
	}
}

func (c *restriction_container) handleInterval(id byte, r restriction, t time.Time) (res bool, time time.Time) {
	time = t
	res = false

	from_time := SetPart(t, r.part, r.from)
	to_time := SetPart(t, r.part, r.to)
	if from_time.Unix() < t.Unix() {
		time = from_time
	}

	for time.Unix() <= to_time.Unix() {
		if temp, temp_time := c.Handle(id, time); temp && (temp_time.Unix() < time.Unix() || !res) {
			time = temp_time
			res = true
		}
	}

	res = res && time.Unix() >= from_time.Unix() && time.Unix() <= to_time.Unix()
	return
}

func (c *restriction_container) handlePeriod(id byte, r restriction, t time.Time) (res bool, time time.Time) {
	v_part := GetPart(t, r.part)
	t = SetPart(t, r.part, v_part+r.value)

	res, time = c.Handle(id, t)

	time_part := GetPart(time, r.part)
	t_part := GetPart(t, r.part)
	res = t_part == time_part

	return
}

func (c *restriction_container) handleExact(id byte, r restriction, t time.Time) (res bool, time time.Time) {
	t = SetPart(t, r.part, r.value)

	res, time = c.Handle(id, t)

	time_part := GetPart(time, r.part)
	t_part := GetPart(t, r.part)
	res = t_part == time_part
	return
}

func GetPart(dt time.Time, p DatePart) byte {
	switch p {
	case DATEPART_YEAR:
		return byte(dt.Year() - 2000)
	case DATEPART_MONTH:
		return byte(dt.Month())
	case DATEPART_DAY:
		return byte(dt.Day())
	case DATEPART_HOUR:
		return byte(dt.Hour())
	case DATEPART_MINUTE:
		return byte(dt.Minute())
	case DATEPART_SECOND:
		return byte(dt.Second())
	default:
		panic("not defined")
	}

	return 0
}

func SetPart(t time.Time, p DatePart, v byte) time.Time {
	switch p {
	case DATEPART_YEAR:
		return time.Date(2000+int(v), 0, 0, 0, 0, 0, 0, nil)
	case DATEPART_MONTH:
		return time.Date(t.Year(), time.Month(v), 0, 0, 0, 0, 0, nil)
	case DATEPART_DAY:
		return time.Date(t.Year(), t.Month(), int(v), 0, 0, 0, 0, nil)
	case DATEPART_HOUR:
		return time.Date(t.Year(), t.Month(), t.Day(), int(v), 0, 0, 0, nil)
	case DATEPART_MINUTE:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), int(v), 0, 0, nil)
	case DATEPART_SECOND:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), int(v), 0, nil)
	default:
		panic(fmt.Sprintf("not implemented DATEPART %v", p))
	}
}

type restriction_builder struct {
	id   byte
	part DatePart
	list *list.List
}

func (b *restriction_builder) Year() {
}

func (b *restriction_builder) Month() {
}

func (b *restriction_builder) Day() {
}

func (b *restriction_builder) Hour() {
}

func (b *restriction_builder) Minute() {
}

func (b *restriction_builder) Second() {
}

func (b *restriction_builder) Root(id byte) {
}

func (b *restriction_builder) Period(period byte) {

}

func (b *restriction_builder) Interval(from byte, to byte) {

}

func (b *restriction_builder) Exact(value byte) {

}
