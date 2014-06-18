package detect

import (
	"testing"
	"time"
)

func Test_Float32(t *testing.T) {

	if v := Float32(0, 1); v != 1 {
		t.Fatalf("%#v", v)
	}

	if v := Float32(); v != 0 {
		t.Fatalf("%#v", v)
	}

}

func Test_Float64(t *testing.T) {

	if v := Float64(0, 1); v != 1 {
		t.Fatalf("%#v", v)
	}

	if v := Float64(); v != 0 {
		t.Fatalf("%#v", v)
	}

}

func Test_Int(t *testing.T) {

	if v := Int(0, 1); v != 1 {
		t.Fatalf("%#v", v)
	}

	if v := Int(); v != 0 {
		t.Fatalf("%#v", v)
	}

}

func Test_Uint(t *testing.T) {

	if v := Uint(0, 1); v != 1 {
		t.Fatalf("%#v", v)
	}

	if v := Uint(); v != 0 {
		t.Fatalf("%#v", v)
	}

}

func Test_String(t *testing.T) {

	if v := String("a", "b"); v != "a" {
		t.Fatalf("%#v", v)
	}

	if v := String(); v != "" {
		t.Fatalf("%#v", v)
	}

}

func Test_Time(t *testing.T) {

	now := time.Now()
	if v := Time(time.Time{}, now); v != now {
		t.Fatalf("%#v", v)
	}

	zero := time.Time{}
	if v := Time(); v != zero {
		t.Fatalf("%#v", v)
	}

}
