package bencoding

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func UnMarshal(src interface{}, dst interface{}) error {
	if src == nil {
		return errors.New("cannot umarshal nil src")
	}

	var rdr *bufio.Reader
	switch t := src.(type) {
	case *bufio.Reader:
		rdr = t
	case io.Reader:
		rdr = bufio.NewReader(t)
	case []byte:
		rdr = bufio.NewReader(bytes.NewReader(t))
	default:
		return errors.New("cannot unmarshal invalid src")
	}
	if dst == nil {
		return errors.New("nil dst")
	}

	var err error
	switch dt := dst.(type) {
	case map[string]interface{}:
		err = readMap(rdr, dt)
	case *[]interface{}:
		err = readSlice(rdr, dt)
	case *string:
		err = readString(rdr, dt)
	case *int:
		err = readInt(rdr, dt)
	default:
		return errors.New("cannot unmarshal invalid dst")
	}

	return err
}

func readMap(rdr *bufio.Reader, dst map[string]interface{}) error {
	return nil
}

func readSlice(rdr *bufio.Reader, dst *[]interface{}) error {
	fb, err := rdr.ReadByte()
	if err != nil {
		return err
	}
	if fb != 'l' {
		return errors.New("not a bencoded list")
	}

	for {

		nb, err := rdr.ReadByte()
		if err != nil {
			return err
		}
		if err := rdr.UnreadByte(); err != nil {
			return err
		}

		switch nb {
		case 'i':
			var val int
			if err := readInt(rdr, &val); err != nil {
				return err
			}
			*dst = append(*dst, val)
		case 'l':
			var val []interface{}
			if err := readSlice(rdr, &val); err != nil {
				return err
			}
			*dst = append(*dst, val)
		case 'd':
			var val map[string]interface{}
			if err := readMap(rdr, val); err != nil {
				return err
			}
			*dst = append(*dst, val)

		case 'e':
			return nil
		default:
			var val string
			if err := readString(rdr, &val); err != nil {
				return err
			}
			*dst = append(*dst, val)
		}
	}
}

func readString(rdr *bufio.Reader, dst *string) error {
	bs, err := rdr.ReadBytes(':')
	if err != nil {
		return err
	}
	len := len(bs)
	strlen := string(bs[:len-1])

	nr, err := strconv.Atoi(strlen)
	if err != nil {
		return err
	}
	strb := []byte{}
	for i := 0; i < nr; i++ {
		by, err := rdr.ReadByte()
		if err != nil {
			return err
		}
		strb = append(strb, by)
	}
	str := string(strb)
	*dst = str
	return nil
}

func readInt(rdr *bufio.Reader, dst *int) error {
	fb, err := rdr.ReadByte()
	if err != nil {
		return err
	}
	if fb != 'i' {
		return errors.New("not a bencoded int")
	}
	bs, err := rdr.ReadBytes('e')
	if err != nil {
		return err
	}
	len := len(bs)
	strlen := string(bs[:len-1])

	nr, err := strconv.Atoi(strlen)
	if err != nil {
		return err
	}

	*dst = nr
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("cannot marshal nil")
	}

	switch t := v.(type) {
	case string:
		return marshalstr(t)
	case int:
		return marshalint(t)
	case []interface{}:
		return marshallist(t)
	case map[string]interface{}:
		return marshaldict(t)
	default:
		return nil, fmt.Errorf("cannot marshal %T", t)
	}

}

func marshalstr(str string) ([]byte, error) {
	res := []byte{}
	res = append(res, strconv.Itoa(len(str))...)
	res = append(res, ':')
	res = append(res, []byte(str)...)

	return res, nil
}

func marshalint(integer int) ([]byte, error) {
	res := []byte{}
	res = append(res, 'i')
	res = append(res, []byte(strconv.Itoa(integer))...)
	res = append(res, 'e')

	return res, nil
}

func marshaldict(dict map[string]interface{}) ([]byte, error) {
	res := []byte{}
	res = append(res, 'd')

	keys := []string{}
	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := dict[key]
		bs, err := marshalstr(key)
		if err != nil {
			return nil, err
		}

		res = append(res, bs...)

		switch t := val.(type) {
		case string:
			bs, err = marshalstr(t)
		case int:
			bs, err = marshalint(t)
		case []interface{}:
			bs, err = marshallist(t)
		case map[string]interface{}:
			bs, err = marshaldict(t)
		}
		if err != nil {
			return nil, err
		}

		res = append(res, bs...)
	}

	res = append(res, 'e')

	return res, nil
}

func marshallist(list []interface{}) ([]byte, error) {
	res := []byte{}
	res = append(res, 'l')

	for _, item := range list {
		var err error
		var bs []byte
		switch t := item.(type) {
		case string:
			bs, err = marshalstr(t)
		case int:
			bs, err = marshalint(t)
		case []interface{}:
			bs, err = marshallist(t)
		case map[string]interface{}:
			bs, err = marshaldict(t)
		}
		if err != nil {
			return nil, err
		}

		res = append(res, bs...)
	}

	res = append(res, 'e')

	return res, nil
}
