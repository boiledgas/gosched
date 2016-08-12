package gosched

import (
	"time"
)

const DATEPART_NOTSET byte = 0xff

const (
	DATEPART_YEAR   byte = 0x01
	DATEPART_MONTH  byte = 0x02
	DATEPART_DAY    byte = 0x03
	DATEPART_HOUR   byte = 0x04
	DATEPART_MINUTE byte = 0x05
	DATEPART_SECOND byte = 0x06
)

// size
type restriction_container struct {
	roots        []byte          // корневые элементы
	restrictions []restriction   // общий массив ограничений строителя
	relations    map[byte][]byte // связи ограничений
	last         time.Time       // последнее время, которое было сгенерировано планировщиком
}

const (
	RESTRICTION_NULL     byte = 0x00
	RESTRICTION_INTERVAL byte = 0x01
	RESTRICTION_EXACT    byte = 0x02
	RESTRICTION_PERIOD   byte = 0x03
)

type restriction struct {
	r_type byte // null, interval, period, time
	part   byte // part of date
	value  byte // period, exact, operator
	from   byte // interval from
	to     byte // interval to
}

type Restriction interface {
	Handle(t time.Time) (time.Time, bool)
}

func NewRestriction() Restriction {
	return &restriction_container{
		restrictions: make([]restriction, 0, 255),
		relations:    make(map[byte][]byte),
		roots:        make([]byte, 0, 255),
	}
}

func (c *restriction_container) Handle(t time.Time) (result time.Time, ok bool) {
	c.last = t
	for _, rootId := range c.roots {
		if temp, handled, _ := c.handle(rootId, t, false); handled {
			if !ok {
				result = temp
				ok = true
			} else {
				if result.Unix() > temp.Unix() {
					result = temp
				}
			}
		}
	}

	return
}

func (c *restriction_container) handle(id byte, t time.Time, generated bool) (result time.Time, ok bool, result_generated bool) {
	ok = false
	result_generated = generated
	restriction := c.restrictions[id]

	switch restriction.r_type {
	case RESTRICTION_INTERVAL:
		{
			if restriction.from != DATEPART_NOTSET {
				from_time := SetPart(t, restriction.part, restriction.from)
				if t.Unix() < from_time.Unix() {
					t = from_time
					result_generated = true
				}
			}
		}
	case RESTRICTION_EXACT:
		{
			t = SetPart(t, restriction.part, restriction.value)
			result_generated = true
		}
	}

handler:
	{
		result = t
		relations_ids, relations_ok := c.relations[id]
		if relations_ok {
			for _, relation_id := range relations_ids {
				if temp_time, temp_ok, temp_generated := c.handle(relation_id, result, result_generated); temp_ok {
					if !ok || result.Unix() > temp_time.Unix() {
						result = temp_time
						result_generated = temp_generated
					}

					ok = true
				}
			}
		} else {
			ok = true
		}
	}

	switch restriction.r_type {
	case RESTRICTION_INTERVAL:
		{
			if ok && restriction.from != DATEPART_NOTSET {
				ok = result.Unix() >= SetPart(t, restriction.part, restriction.from).Unix()
			}
			if ok && restriction.to != DATEPART_NOTSET {
				ok = result.Unix() <= SetPart(t, restriction.part, restriction.to).Unix()
			}
		}
	case RESTRICTION_PERIOD:
		{
			if !result_generated {
				t = SetPart(t, restriction.part, GetPart(t, restriction.part)+restriction.value)
				result_generated = true
				goto handler
			}

			temp_part := GetPart(t, restriction.part)
			new_part := GetPart(result, restriction.part)
			ok = temp_part == new_part
		}
	case RESTRICTION_EXACT:
		{
			temp_part := GetPart(t, restriction.part)
			new_part := GetPart(result, restriction.part)
			ok = temp_part == new_part && c.last.Unix() != result.Unix()
		}
	}

	// если возвращается false, дата почемуто 0001-01-01
	return
}

func GetPart(dt time.Time, p byte) (result byte) {
	switch p {
	case DATEPART_YEAR:
		result = byte(dt.Year() - 2000)
	case DATEPART_MONTH:
		result = byte(dt.Month())
	case DATEPART_DAY:
		result = byte(dt.Day())
	case DATEPART_HOUR:
		result = byte(dt.Hour())
	case DATEPART_MINUTE:
		result = byte(dt.Minute())
	case DATEPART_SECOND:
		result = byte(dt.Second())
	}
	return
}

func SetPart(t time.Time, p byte, v byte) (result time.Time) {
	switch p {
	case DATEPART_YEAR:
		result = time.Date(2000+int(v), 1, 1, 0, 0, 0, 0, t.Location())
	case DATEPART_MONTH:
		result = time.Date(t.Year(), time.Month(v), 1, 0, 0, 0, 0, t.Location())
	case DATEPART_DAY:
		result = time.Date(t.Year(), t.Month(), int(v), 0, 0, 0, 0, t.Location())
	case DATEPART_HOUR:
		result = time.Date(t.Year(), t.Month(), t.Day(), int(v), 0, 0, 0, t.Location())
	case DATEPART_MINUTE:
		result = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), int(v), 0, 0, t.Location())
	case DATEPART_SECOND:
		result = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), int(v), 0, t.Location())
	}
	return
}
