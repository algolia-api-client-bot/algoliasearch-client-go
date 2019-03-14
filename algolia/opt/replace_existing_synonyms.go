// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ReplaceExistingSynonymsOption struct {
	value bool
}

func ReplaceExistingSynonyms(v bool) *ReplaceExistingSynonymsOption {
	return &ReplaceExistingSynonymsOption{v}
}

func (o ReplaceExistingSynonymsOption) Get() bool {
	return o.value
}

func (o ReplaceExistingSynonymsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ReplaceExistingSynonymsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *ReplaceExistingSynonymsOption) Equal(o2 *ReplaceExistingSynonymsOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func ReplaceExistingSynonymsEqual(o1, o2 *ReplaceExistingSynonymsOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
