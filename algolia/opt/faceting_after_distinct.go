// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type FacetingAfterDistinctOption struct {
	value bool
}

func FacetingAfterDistinct(v bool) *FacetingAfterDistinctOption {
	return &FacetingAfterDistinctOption{v}
}

func (o FacetingAfterDistinctOption) Get() bool {
	return o.value
}

func (o FacetingAfterDistinctOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *FacetingAfterDistinctOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *FacetingAfterDistinctOption) Equal(o2 *FacetingAfterDistinctOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func FacetingAfterDistinctEqual(o1, o2 *FacetingAfterDistinctOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
