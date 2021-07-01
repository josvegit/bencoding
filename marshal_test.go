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
