// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type PercentileComputationOption struct {
	value bool
}

func PercentileComputation(v bool) *PercentileComputationOption {
	return &PercentileComputationOption{v}
}

func (o PercentileComputationOption) Get() bool {
	return o.value
}

func (o PercentileComputationOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *PercentileComputationOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = true
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *PercentileComputationOption) Equal(o2 *PercentileComputationOption) bool {
	if o2 == nil {
		return o.value == true
	}
	return o.value == o2.value
}

func PercentileComputationEqual(o1, o2 *PercentileComputationOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
