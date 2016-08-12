package gosched

type year_restriction_builder interface {
	YearBetween(byte, byte) RestrictionBuilder
	YearFrom(byte) RestrictionBuilder
	YearTo(byte) RestrictionBuilder
	YearEvery(byte) RestrictionBuilder
	Year(byte) RestrictionBuilder

	YearBetweenAll(byte, byte, BuilderFunc)
	YearFromAll(byte, BuilderFunc)
	YearToAll(byte, BuilderFunc)
	YearEveryAll(byte, BuilderFunc)
	YearAll(byte, BuilderFunc)
}
type month_restriction_builder interface {
	MonthBetween(byte, byte) RestrictionBuilder
	MonthFrom(byte) RestrictionBuilder
	MonthTo(byte) RestrictionBuilder
	MonthEvery(byte) RestrictionBuilder
	Month(byte) RestrictionBuilder

	MonthBetweenAll(byte, byte, BuilderFunc)
	MonthFromAll(byte, BuilderFunc)
	MonthToAll(byte, BuilderFunc)
	MonthEveryAll(byte, BuilderFunc)
	MonthAll(byte, BuilderFunc)
}
type day_restriction_builder interface {
	DayBetween(byte, byte) RestrictionBuilder
	DayFrom(byte) RestrictionBuilder
	DayTo(byte) RestrictionBuilder
	DayEvery(byte) RestrictionBuilder
	Day(byte) RestrictionBuilder

	DayBetweenAll(byte, byte, BuilderFunc)
	DayFromAll(byte, BuilderFunc)
	DayToAll(byte, BuilderFunc)
	DayEveryAll(byte, BuilderFunc)
	DayAll(byte, BuilderFunc)
}
type hour_restriction_builder interface {
	HourBetween(byte, byte) RestrictionBuilder
	HourFrom(byte) RestrictionBuilder
	HourTo(byte) RestrictionBuilder
	HourEvery(byte) RestrictionBuilder
	Hour(byte) RestrictionBuilder

	HourBetweenAll(byte, byte, BuilderFunc)
	HourFromAll(byte, BuilderFunc)
	HourToAll(byte, BuilderFunc)
	HourEveryAll(byte, BuilderFunc)
	HourAll(byte, BuilderFunc)
}
type minute_restriction_builder interface {
	MinuteBetween(byte, byte) RestrictionBuilder
	MinuteFrom(byte) RestrictionBuilder
	MinuteTo(byte) RestrictionBuilder
	MinuteEvery(byte) RestrictionBuilder
	Minute(byte) RestrictionBuilder

	MinuteBetweenAll(byte, byte, BuilderFunc)
	MinuteFromAll(byte, BuilderFunc)
	MinuteToAll(byte, BuilderFunc)
	MinuteEveryAll(byte, BuilderFunc)
	MinuteAll(byte, BuilderFunc)
}
type second_restriction_builder interface {
	SecondBetween(byte, byte) RestrictionBuilder
	SecondFrom(byte) RestrictionBuilder
	SecondTo(byte) RestrictionBuilder
	SecondEvery(byte) RestrictionBuilder
	Second(byte) RestrictionBuilder

	SecondBetweenAll(byte, byte, BuilderFunc)
	SecondFromAll(byte, BuilderFunc)
	SecondToAll(byte, BuilderFunc)
	SecondEveryAll(byte, BuilderFunc)
	SecondAll(byte, BuilderFunc)
}

func (b restriction_builder) YearBetween(from byte, to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_YEAR, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) YearFrom(from byte) RestrictionBuilder {
	r := restriction{part: DATEPART_YEAR, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) YearTo(to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_YEAR, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) YearEvery(period byte) RestrictionBuilder {
	r := restriction{part: DATEPART_YEAR, value: period, r_type: RESTRICTION_PERIOD}
	return b.restriction(r)
}

func (b restriction_builder) Year(year byte) RestrictionBuilder {
	r := restriction{part: DATEPART_YEAR, value: year, r_type: RESTRICTION_EXACT}
	return b.restriction(r)
}

func (b restriction_builder) YearBetweenAll(from byte, to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_YEAR, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) YearFromAll(from byte, f BuilderFunc) {
	r := restriction{part: DATEPART_YEAR, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) YearToAll(to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_YEAR, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) YearEveryAll(period byte, f BuilderFunc) {
	r := restriction{part: DATEPART_YEAR, value: period, r_type: RESTRICTION_PERIOD}
	b.all_restriction(r, f)
}

func (b restriction_builder) YearAll(year byte, f BuilderFunc) {
	r := restriction{part: DATEPART_YEAR, value: year, r_type: RESTRICTION_EXACT}
	b.all_restriction(r, f)
}

func (b restriction_builder) MonthBetween(from byte, to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MONTH, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) MonthFrom(from byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MONTH, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) MonthTo(to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MONTH, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) MonthEvery(period byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MONTH, value: period, r_type: RESTRICTION_PERIOD}
	return b.restriction(r)
}

func (b restriction_builder) Month(year byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MONTH, value: year, r_type: RESTRICTION_EXACT}
	return b.restriction(r)
}

func (b restriction_builder) MonthBetweenAll(from byte, to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MONTH, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) MonthFromAll(from byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MONTH, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) MonthToAll(to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MONTH, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) MonthEveryAll(period byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MONTH, value: period, r_type: RESTRICTION_PERIOD}
	b.all_restriction(r, f)
}

func (b restriction_builder) MonthAll(year byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MONTH, value: year, r_type: RESTRICTION_EXACT}
	b.all_restriction(r, f)
}

func (b restriction_builder) DayBetween(from byte, to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_DAY, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) DayFrom(from byte) RestrictionBuilder {
	r := restriction{part: DATEPART_DAY, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) DayTo(to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_DAY, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) DayEvery(period byte) RestrictionBuilder {
	r := restriction{part: DATEPART_DAY, value: period, r_type: RESTRICTION_PERIOD}
	return b.restriction(r)
}

func (b restriction_builder) Day(year byte) RestrictionBuilder {
	r := restriction{part: DATEPART_DAY, value: year, r_type: RESTRICTION_EXACT}
	return b.restriction(r)
}

func (b restriction_builder) DayBetweenAll(from byte, to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_DAY, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) DayFromAll(from byte, f BuilderFunc) {
	r := restriction{part: DATEPART_DAY, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) DayToAll(to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_DAY, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) DayEveryAll(period byte, f BuilderFunc) {
	r := restriction{part: DATEPART_DAY, value: period, r_type: RESTRICTION_PERIOD}
	b.all_restriction(r, f)
}

func (b restriction_builder) DayAll(year byte, f BuilderFunc) {
	r := restriction{part: DATEPART_DAY, value: year, r_type: RESTRICTION_EXACT}
	b.all_restriction(r, f)
}

func (b restriction_builder) HourBetween(from byte, to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_HOUR, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) HourFrom(from byte) RestrictionBuilder {
	r := restriction{part: DATEPART_HOUR, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) HourTo(to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_HOUR, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) HourEvery(period byte) RestrictionBuilder {
	r := restriction{part: DATEPART_HOUR, value: period, r_type: RESTRICTION_PERIOD}
	return b.restriction(r)
}

func (b restriction_builder) Hour(year byte) RestrictionBuilder {
	r := restriction{part: DATEPART_HOUR, value: year, r_type: RESTRICTION_EXACT}
	return b.restriction(r)
}

func (b restriction_builder) HourBetweenAll(from byte, to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_HOUR, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) HourFromAll(from byte, f BuilderFunc) {
	r := restriction{part: DATEPART_HOUR, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) HourToAll(to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_HOUR, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) HourEveryAll(period byte, f BuilderFunc) {
	r := restriction{part: DATEPART_HOUR, value: period, r_type: RESTRICTION_PERIOD}
	b.all_restriction(r, f)
}

func (b restriction_builder) HourAll(year byte, f BuilderFunc) {
	r := restriction{part: DATEPART_HOUR, value: year, r_type: RESTRICTION_EXACT}
	b.all_restriction(r, f)
}

func (b restriction_builder) MinuteBetween(from byte, to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MINUTE, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) MinuteFrom(from byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MINUTE, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) MinuteTo(to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MINUTE, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) MinuteEvery(period byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MINUTE, value: period, r_type: RESTRICTION_PERIOD}
	return b.restriction(r)
}

func (b restriction_builder) Minute(year byte) RestrictionBuilder {
	r := restriction{part: DATEPART_MINUTE, value: year, r_type: RESTRICTION_EXACT}
	return b.restriction(r)
}

func (b restriction_builder) MinuteBetweenAll(from byte, to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MINUTE, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) MinuteFromAll(from byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MINUTE, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) MinuteToAll(to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MINUTE, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) MinuteEveryAll(period byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MINUTE, value: period, r_type: RESTRICTION_PERIOD}
	b.all_restriction(r, f)
}

func (b restriction_builder) MinuteAll(year byte, f BuilderFunc) {
	r := restriction{part: DATEPART_MINUTE, value: year, r_type: RESTRICTION_EXACT}
	b.all_restriction(r, f)
}

func (b restriction_builder) SecondBetween(from byte, to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_SECOND, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) SecondFrom(from byte) RestrictionBuilder {
	r := restriction{part: DATEPART_SECOND, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) SecondTo(to byte) RestrictionBuilder {
	r := restriction{part: DATEPART_SECOND, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	return b.restriction(r)
}

func (b restriction_builder) SecondEvery(period byte) RestrictionBuilder {
	r := restriction{part: DATEPART_SECOND, value: period, r_type: RESTRICTION_PERIOD}
	return b.restriction(r)
}

func (b restriction_builder) Second(year byte) RestrictionBuilder {
	r := restriction{part: DATEPART_SECOND, value: year, r_type: RESTRICTION_EXACT}
	return b.restriction(r)
}

func (b restriction_builder) SecondBetweenAll(from byte, to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_SECOND, from: from, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) SecondFromAll(from byte, f BuilderFunc) {
	r := restriction{part: DATEPART_SECOND, from: from, to: DATEPART_NOTSET, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) SecondToAll(to byte, f BuilderFunc) {
	r := restriction{part: DATEPART_SECOND, from: DATEPART_NOTSET, to: to, r_type: RESTRICTION_INTERVAL}
	b.all_restriction(r, f)
}

func (b restriction_builder) SecondEveryAll(period byte, f BuilderFunc) {
	r := restriction{part: DATEPART_SECOND, value: period, r_type: RESTRICTION_PERIOD}
	b.all_restriction(r, f)
}

func (b restriction_builder) SecondAll(year byte, f BuilderFunc) {
	r := restriction{part: DATEPART_SECOND, value: year, r_type: RESTRICTION_EXACT}
	b.all_restriction(r, f)
}
