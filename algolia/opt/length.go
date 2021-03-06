// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// LengthOption is a wrapper for an Length option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type LengthOption struct {
	value int
}

// Length wraps the given value into a LengthOption.
func Length(v int) *LengthOption {
	return &LengthOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *LengthOption) Get() int {
	if o == nil {
		return 0
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// LengthOption.
func (o LengthOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// LengthOption.
func (o *LengthOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *LengthOption) Equal(o2 *LengthOption) bool {
	if o2 == nil {
		return o.value == 0
	}
	return o.value == o2.value
}

// LengthEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func LengthEqual(o1, o2 *LengthOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
