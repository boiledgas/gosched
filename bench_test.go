package gosched

import (
	"testing"
	"time"
)

func pointer_receiver(t *time.Time, count int) {
	r := (*t)
	for i := 0; i < count; i++ {
		r = r.Add(time.Second * 40)
	}

	(*t) = r
}

func return_value(t time.Time, count int) time.Time {
	for i := 0; i < count; i++ {
		t = t.Add(time.Second * 40)
	}

	return t
}

var count int = 1

// conclusion:
//	- use copy value to set time.Time
//	- use pointer value to get values on time.Time
//	- use copy value is faster than get pointer on time.Time

func Benchmark_PointerReceiver_New(b *testing.B) {
	t := time.Now()
	for i := 0; i < b.N; i++ {
		new_t := t
		pointer_receiver(&new_t, count)
		if new_t.Unix() != t.Unix()+int64(count*40) {
			panic("math")
		}
	}
}

func Benchmark_ReturnValue_New(b *testing.B) {
	t := time.Now()
	for i := 0; i < b.N; i++ {
		new_t := return_value(t, count)
		if new_t.Unix() != t.Unix()+int64(count*40) {
			panic("math")
		}
	}
}

func Benchmark_PointerReceiver_Self(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		t = time.Now()
		pointer_receiver(&t, count)
	}
}

func Benchmark_ReturnValue_Self(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		t = time.Now()
		t = return_value(t, count)
	}
}

func Benchmark_PointerReceiver_SelfSingle(b *testing.B) {
	t := time.Now()
	for i := 0; i < b.N; i++ {
		pointer_receiver(&t, count)
	}
}

func Benchmark_ReturnValue_SelfSingle(b *testing.B) {
	t := time.Now()
	for i := 0; i < b.N; i++ {
		t = return_value(t, count)
	}
}

func get_seconds_value(t time.Time) int {
	return t.Second()
}

func get_seconds_pointer(t *time.Time) int {
	return t.Second()
}

func Benchmark_GetValue_Copy(b *testing.B) {
	t := time.Now()
	s := 0
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			s = get_seconds_value(t)
			if s < 0 {
				panic("s < 0")
			}
		}
	}
}

func Benchmark_GetValue_Pointer(b *testing.B) {
	t := time.Now()
	s := 0
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			s = get_seconds_pointer(&t)
			if s < 0 {
				panic("s < 0")
			}
		}
	}
}

func empty_value(t time.Time) {
}

func pointer_value(t *time.Time) {
}

func Benchmark_EmptyValue_Copy(b *testing.B) {
	t := time.Now()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			empty_value(t)
		}
	}
}

func Benchmark_EmptyValue_Pointer(b *testing.B) {
	t := time.Now()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; j++ {
			pointer_value(&t)
		}
	}
}
