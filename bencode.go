package bencoding

import (
	"errors"
	"fmt"
	"strconv"
)

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
	return nil, nil
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
