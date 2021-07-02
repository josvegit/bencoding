package bencoding

import (
	"testing"
)

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

func Test_UnMarshal_Invalid_Dst(t *testing.T) {
	type test struct {
		input interface{}
	}

	strdst := ""

	if err := UnMarshal([]byte("3:cat"), strdst); err == nil {
		t.Fatal("cannot unmarshal into string")
	}
}

func Test_UnMarshal_String_Dst(t *testing.T) {
	strdst := ""

	if err := UnMarshal([]byte("3:cat"), &strdst); err != nil {
		t.Fatal(err.Error())
	}

	if strdst != "cat" {
		t.Fatalf("wanted cat got %s", strdst)
	}
}

func Test_UnMarshal_Int_Dst(t *testing.T) {
	intdst := 0

	if err := UnMarshal([]byte("i12222222e"), &intdst); err != nil {
		t.Fatal(err.Error())
	}

	if intdst != 12222222 {
		t.Fatalf("wanted 12222222 got %d", intdst)
	}
}

func Test_UnMarshal_List_Dst(t *testing.T) {
	listdst := []interface{}{}

	if err := UnMarshal([]byte("l4:coole"), &listdst); err != nil {
		t.Fatal(err.Error())
	}

	if listdst[0] != "cool" {
		t.Fatal("invalid marshalled list")
	}
}

func Test_UnMarshal_Dict_Dst(t *testing.T) {
	dict := map[string]interface{}{
		"info": []interface{}{
			map[string]interface{}{
				"id":    1,
				"ip":    "192.168.0.1",
				"peers": []interface{}{"a", "b", "c"},
			},
			map[string]interface{}{
				"id":    2,
				"ip":    "192.168.0.2",
				"peers": []interface{}{"d", "e", "f"},
			},
			map[string]interface{}{
				"id":    3,
				"ip":    "192.168.0.3",
				"peers": []interface{}{"g", "h", "i"},
			},
			map[string]interface{}{
				"id":    4,
				"ip":    "192.168.0.4",
				"peers": []interface{}{"j", "k", "l"},
			},
			map[string]interface{}{
				"id":    5,
				"ip":    "192.168.0.5",
				"peers": []interface{}{"m", "n", "o"},
			},
		},
	}

	bs, err := Marshal(dict)
	if err != nil {
		t.Fatal(err.Error())
	}

	dst := make(map[string]interface{})
	if err := UnMarshal(bs, dst); err != nil {
		t.Fatal(err.Error())
	}

	var list []interface{}
	switch r := dst["into"].(type) {
	case []interface{}:
		list = r
	case nil:
		t.Fatal("dst was nil")
	default:
		t.Fatal("invalid type for dst")
	}

	if len(list) != 5 {
		t.Fatal("invalid length of list")
	}

	first := list[0].(map[string]interface{})
	last := list[4].(map[string]interface{})

	if first["id"] != 1 {
		t.Fatal("invalid result dict")
	}

	if last["ip"] != "192.168.0.5" {
		t.Fatal("invalid result dict")
	}

	lastIps := last["peers"].([]interface{})
	if len(lastIps) != 3 {
		t.Fatal("invalid result dict")
	}

	peer := lastIps[0].(string)

	if peer != "m" {
		t.Fatal("invalid result dict")
	}

}
