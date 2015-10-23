package decoder

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"reflect"
	"strconv"
)

// Decoder decodes values from a map[string][]string to a struct.
type Decoder struct {
}

// NewDecoder = decoder constructor
func NewDecoder() *Decoder {
	return &Decoder{}
}

// DecodeReqBody decode the request body to a struct
func (d *Decoder) DecodeReqBody(dst interface{}, src io.ReadCloser) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("schema: interface must be a pointer to struct")
	}
	return doDecodeReqBody(dst, src)
}

// DecodeURLValues  decode the url values to a struct
func (d *Decoder) DecodeURLValues(dst interface{}, src url.Values) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("schema: interface must be a pointer to struct")
	}
	return doDecodeURLValues(dst, src)
}

func doDecodeReqBody(dst interface{}, src io.ReadCloser) error {
	decoder := json.NewDecoder(src)
	return decoder.Decode(dst)
}

func doDecodeURLValues(_struct interface{}, query url.Values) error {
	newInstance := reflect.ValueOf(_struct).Elem()

	for k, v := range query {
		field := newInstance.FieldByName(k)

		switch field.Kind() {
		case reflect.Int:
			intValue, err := strconv.ParseInt(v[0], 0, 64)
			if err != nil {
				return err
			}

			field.SetInt(intValue)

		case reflect.String:
			field.SetString(v[0])
			return nil
		}

	}
	return nil
}
