// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"reflect"
)

type DisablePrefixOnAttributesOption struct {
	value []string
}

func DisablePrefixOnAttributes(v ...string) *DisablePrefixOnAttributesOption {
	return &DisablePrefixOnAttributesOption{v}
}

func (o DisablePrefixOnAttributesOption) Get() []string {
	return o.value
}

func (o DisablePrefixOnAttributesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *DisablePrefixOnAttributesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = []string{}
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *DisablePrefixOnAttributesOption) Equal(o2 *DisablePrefixOnAttributesOption) bool {
	if o2 == nil {
		return reflect.DeepEqual(o.value, []string{})
	}
	return reflect.DeepEqual(o.value, o2.value)
}

func DisablePrefixOnAttributesEqual(o1, o2 *DisablePrefixOnAttributesOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
