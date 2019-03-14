// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type PageOption struct {
	value int
}

func Page(v int) *PageOption {
	return &PageOption{v}
}

func (o PageOption) Get() int {
	return o.value
}

func (o PageOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *PageOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *PageOption) Equal(o2 *PageOption) bool {
	if o2 == nil {
		return o.value == 0
	}
	return o.value == o2.value
}

func PageEqual(o1, o2 *PageOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
