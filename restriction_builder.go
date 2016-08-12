package gosched

import "fmt"

type BuilderFunc func(b RestrictionBuilder)
type base_restriction_builder interface {
	restriction(r restriction) RestrictionBuilder
	all_restriction(r restriction, f BuilderFunc)
}

type RestrictionBuilder interface {
	base_restriction_builder
	year_restriction_builder
	month_restriction_builder
	day_restriction_builder
	hour_restriction_builder
	minute_restriction_builder
	second_restriction_builder
}

type restriction_builder struct {
	id     byte // идентификатор ограничения
	part   byte // часть ограничения
	r_type byte // тип ограничения

	C *restriction_container
}

func NewBuilder(r Restriction) RestrictionBuilder {
	return restriction_builder{
		id: DATEPART_NOTSET,
		C:  r.(*restriction_container),
	}
}

func (b restriction_builder) restriction(r restriction) (builder RestrictionBuilder) {
	if b.id != DATEPART_NOTSET {
		if b.r_type == RESTRICTION_INTERVAL && r.r_type != RESTRICTION_INTERVAL {
			if r.part < b.part {
				panic(fmt.Sprintf("builder datepart %v less than %v", b.part, r.part))
			}
		} else {
			if r.part <= b.part {
				panic(fmt.Sprintf("builder(%v) part %v is less or equal than restriction(%v) %v", b.r_type, b.part, r.r_type, r.part))
			}
		}
	}

	b.C.restrictions = append(b.C.restrictions, r)
	id := byte(len(b.C.restrictions) - 1)

	if b.id == DATEPART_NOTSET {
		b.C.roots = append(b.C.roots, id)
	} else {
		if _, ok := b.C.relations[b.id]; !ok {
			b.C.relations[b.id] = make([]byte, 0, 3)
		}
		b.C.relations[b.id] = append(b.C.relations[b.id], id)
	}

	return restriction_builder{id: id, r_type: r.r_type, part: r.part, C: b.C}
}

func (b restriction_builder) all_restriction(r restriction, f BuilderFunc) {
	if b.id != DATEPART_NOTSET && byte(r.part) <= byte(b.part) {
		panic(fmt.Sprintf("root datepart %v is less than %v", b.part, r.part))
	}

	b.C.restrictions = append(b.C.restrictions, r)
	id := byte(len(b.C.restrictions) - 1)

	if b.id == DATEPART_NOTSET {
		b.C.roots = append(b.C.roots, id)
	} else {
		if _, ok := b.C.relations[b.id]; !ok {
			b.C.relations[b.id] = make([]byte, 0, 3)
		}
		b.C.relations[b.id] = append(b.C.relations[b.id], id)
	}

	f(restriction_builder{id: id, r_type: r.r_type, part: r.part, C: b.C})
}
