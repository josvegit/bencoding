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
