package bencoding

import "testing"

func Test_Marshal_Nil(t *testing.T) {
	_, err := Marshal(nil)
	if err == nil {
		t.Fatal("nil value should produce an error")
	}
}

func Test_Marshal_Invalid(t *testing.T) {
	_, err := Marshal([]byte("asasdsad"))
	if err == nil {
		t.Fatal("a byte slice is not something we can marshal")
	}
}

func Test_Marshal_Valid(t *testing.T) {
	type test struct {
		input interface{}
		want  string
	}

	table := []test{
		{
			input: "",
			want:  "0:",
		},
		{
			input: "cool",
			want:  "4:cool",
		},
		{
			input: 0,
			want:  "i0e",
		},
		{
			input: 123,
			want:  "i123e",
		},
		{
			input: []interface{}{},
			want:  "le",
		},
		{
			input: []interface{}{
				"cool", "cat", 123, []interface{}{
					map[string]interface{}{
						"zool": "cat",
						"cat":  "cool",
						"vat":  "cool",
					},
				},
			},
			want: "l4:cool3:cati123eld3:cat4:cool3:vat4:cool4:zool3:cateee",
		},
	}
	for _, test := range table {
		got, err := Marshal(test.input)
		if err != nil {
			t.Fatal(err.Error())
		}

		if string(got) != test.want {
			t.Fatalf("got %s, wanted %s", string(got), test.want)
		}
	}
}

func Test_Marshal_Deep(t *testing.T) {
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
	got := string(bs)

	if got != "d4:infold2:idi1e2:ip11:192.168.0.15:peersl1:a1:b1:ceed2:idi2e2:ip11:192.168.0.25:peersl1:d1:e1:feed2:idi3e2:ip11:192.168.0.35:peersl1:g1:h1:ieed2:idi4e2:ip11:192.168.0.45:peersl1:j1:k1:leed2:idi5e2:ip11:192.168.0.55:peersl1:m1:n1:oeeee" {
		t.Fatalf("wanted d4:infold2:idi1e2:ip11:192.168.0.15:peersl1:a1:b1:ceed2:idi2e2:ip11:192.168.0.25:peersl1:d1:e1:feed2:idi3e2:ip11:192.168.0.35:peersl1:g1:h1:ieed2:idi4e2:ip11:192.168.0.45:peersl1:j1:k1:leed2:idi5e2:ip11:192.168.0.55:peersl1:m1:n1:oeeee got %s", got)
	}

}
