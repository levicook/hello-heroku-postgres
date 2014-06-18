package detect

import "time"

func Float32(floats ...float32) float32 {
	for _, f := range floats {
		if f != 0 {
			return f
		}
	}
	return 0
}

func Float64(floats ...float64) float64 {
	for _, f := range floats {
		if f != 0 {
			return f
		}
	}
	return 0
}

func Int(ints ...int) int {
	for _, v := range ints {
		if v != 0 {
			return v
		}
	}
	return 0
}

func Uint(uints ...uint) uint {
	for _, v := range uints {
		if v != 0 {
			return v
		}
	}
	return 0
}

func String(strings ...string) string {
	for _, s := range strings {
		if s != "" {
			return s
		}
	}
	return ""
}

func Time(times ...time.Time) time.Time {
	for _, t := range times {
		if !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}
