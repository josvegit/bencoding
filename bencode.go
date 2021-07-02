package bencoding

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
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
		err = unMarshalMap(rdr, dt)
	case []interface{}:
		err = unMarshalSlice(rdr, dt)
	case *string:
		err = unMarshalString(rdr, dt)
	case *int:
		err = unMarshalInt(rdr, dt)
	default:
		return errors.New("cannot unmarshal invalid dst")
	}

	return err
}

func unMarshalMap(rdr *bufio.Reader, dst map[string]interface{}) error {
	return nil
}

func unMarshalSlice(rdr *bufio.Reader, dst []interface{}) error {
	return nil
}

func unMarshalString(rdr *bufio.Reader, dst *string) error {
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

func unMarshalInt(rdr *bufio.Reader, dst *int) error {
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
