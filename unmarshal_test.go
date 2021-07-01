package bencoding

import "testing"

func Test_UnMarshal_Nil_Src(t *testing.T) {
	if err := UnMarshal(nil, nil); err == nil {
		t.Fatal("cannot unmarshal nil src")
	}
}

func Test_Unmarshal_Invalid_Src(t *testing.T) {
	if err := UnMarshal("string", nil); err == nil {
		t.Fatal("cannot unmarshal invalid src")
	}
}

func Test_UnMarshal_Nil_Dst(t *testing.T) {
	if err := UnMarshal([]byte("asds"), nil); err == nil {
		t.Fatal("cannot unmarshal nil dst")
	}
}

func Test_UnMarshal_NoPointer_Dst(t *testing.T) {
	if err := UnMarshal([]byte("asds"), "str"); err == nil {
		t.Fatal("cannot unmarshal nonpointer dst")
	}
}
