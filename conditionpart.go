package gosched

import (
	"fmt"
	"time"
)

type ConditionPart byte

const (
	CONDITION_PART_YEAR   ConditionPart = 0x01
	CONDITION_PART_MONTH  ConditionPart = 0x02
	CONDITION_PART_DAY    ConditionPart = 0x03
	CONDITION_PART_HOUR   ConditionPart = 0x04
	CONDITION_PART_MIN ConditionPart = 0x05
	CONDITION_PART_SEC ConditionPart = 0x06
)

func (p ConditionPart) Get(dt time.Time) byte {
	switch p {
	case CONDITION_PART_YEAR:
		return byte(dt.Year() - 2000)
	case CONDITION_PART_MONTH:
		return byte(dt.Month())
	case CONDITION_PART_DAY:
		return byte(dt.Day())
	case CONDITION_PART_HOUR:
		return byte(dt.Hour())
	case CONDITION_PART_MIN:
		return byte(dt.Minute())
	case CONDITION_PART_SEC:
		return byte(dt.Second())
	default:
		panic("not defined")
	}

	return 0
}

func (p ConditionPart) Set(t time.Time, v byte) time.Time {
	switch p {
	case CONDITION_PART_YEAR:
		return time.Date(2000+int(v), 1, 1, 0, 0, 0, 0, t.Location())
	case CONDITION_PART_MONTH:
		return time.Date(t.Year(), time.Month(v), 1, 0, 0, 0, 0, t.Location())
	case CONDITION_PART_DAY:
		return time.Date(t.Year(), t.Month(), int(v), 0, 0, 0, 0, t.Location())
	case CONDITION_PART_HOUR:
		return time.Date(t.Year(), t.Month(), t.Day(), int(v), 0, 0, 0, t.Location())
	case CONDITION_PART_MIN:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), int(v), 0, 0, t.Location())
	case CONDITION_PART_SEC:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), int(v), 0, t.Location())
	default:
		panic(fmt.Sprintf("not implemented DATEPART %v", p))
	}
}
