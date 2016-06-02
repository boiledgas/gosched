package gosched

import "time"

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

// size 48 (but allocations on heap)
type restriction_container struct {
	restrictions []restriction
	relations    [][]byte
	rootId       byte
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

func (c *restriction_container) Handle(restrictionId byte, t time.Time) time.Time {
	r := c.restrictions[restrictionId]
	r.Handle(t)

	relations := c.relations[restrictionId]

	t_min := int64(0xFFFFFFFF)
	for _, nodeId := range relations {
		r := c.restrictions[nodeId]
		if t_next := r.Handle(t); t_min < t_next.Unix() {
			t_min = t_next.Unix()
		}
	}
	return time.Unix(t_min, 0)
}

func (r *restriction) Handle(t time.Time) time.Time {
	switch r.r_type {
	case TIMER_INTERVAL:
		return r.handleInterval(t)
	case TIMER_PERIOD:
		return r.handlePeriod(t)
	case TIMER_TIME:
		return r.handleTime(t)
	default:
		panic("unknown r_type")
	}
}

func (r *restriction) handleInterval(t time.Time) time.Time {

	return t
}

func (r *restriction) handlePeriod(t time.Time) time.Time {
	return t
}

func (r *restriction) handleTime(t time.Time) time.Time {
	return t
}

func getPart(p DatePart, t time.Time) byte {
	switch p {
	case DATEPART_YEAR:
		return byte(t.Year() - 2000)
	case DATEPART_MONTH:
		return byte(t.Month())
	case DATEPART_DAY:
		return byte(t.Day())
	case DATEPART_HOUR:
		return byte(t.Hour())
	case DATEPART_MINUTE:
		return byte(t.Minute())
	case DATEPART_SECOND:
		return byte(t.Second())
	default:
		panic("not defined")
	}

	return 0
}

func setPart(t *time.Time, p DatePart, v byte) {
	switch p {
	case DATEPART_YEAR:
		(*t) = time.Date(2000+int(v), 0, 0, 0, 0, 0, 0, nil)
	case DATEPART_MONTH:
		(*t) = time.Date(t.Year(), time.Month(v), 0, 0, 0, 0, 0, nil)
	case DATEPART_DAY:
		(*t) = time.Date(t.Year(), t.Month(), int(v), 0, 0, 0, 0, nil)
	case DATEPART_HOUR:
		(*t) = time.Date(t.Year(), t.Month(), t.Day(), int(v), 0, 0, 0, nil)
	case DATEPART_MINUTE:
		(*t) = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), int(v), 0, 0, nil)
	case DATEPART_SECOND:
		(*t) = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), int(v), 0, nil)
	default:
		panic("not implemented DATEPART")
	}
}
