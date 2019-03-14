// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type LengthOption struct {
	value int
}

func Length(v int) *LengthOption {
	return &LengthOption{v}
}

func (o LengthOption) Get() int {
	return o.value
}

func (o LengthOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *LengthOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *LengthOption) Equal(o2 *LengthOption) bool {
	if o2 == nil {
		return o.value == 0
	}
	return o.value == o2.value
}

func LengthEqual(o1, o2 *LengthOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
