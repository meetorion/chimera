//go:build !(linux || windows || darwin)

package jsonKit

import "encoding/json"

type impl struct {
}

func (i *impl) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (i *impl) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (i *impl) MarshalToString(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	return string(data), err
}

func (i *impl) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (i *impl) UnmarshalFromString(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}

func init() {
	library = "encoding/json"

	i := &impl{}
	defaultApi = i
	stdApi = i
}
