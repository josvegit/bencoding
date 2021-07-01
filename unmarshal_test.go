package bencoding

import "testing"

func Test_UnMarshal_Nil_In(t *testing.T) {
	err := UnMarshal(nil, nil)
	if err == nil {
		t.Fatal("cannot unmarshal into nil")
	}
}

func Test_Unmarshal_NonPointer_In(t *testing.T) {
	err := UnMarshal("string", nil)
	if err == nil {
		t.Fatal("cannot unmarshal into non pointer")
	}
}
