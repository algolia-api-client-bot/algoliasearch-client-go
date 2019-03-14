// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ClearExistingRulesOption struct {
	value bool
}

func ClearExistingRules(v bool) *ClearExistingRulesOption {
	return &ClearExistingRulesOption{v}
}

func (o ClearExistingRulesOption) Get() bool {
	return o.value
}

func (o ClearExistingRulesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ClearExistingRulesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *ClearExistingRulesOption) Equal(o2 *ClearExistingRulesOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func ClearExistingRulesEqual(o1, o2 *ClearExistingRulesOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
