// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

type ForwardToReplicasOption struct {
	value bool
}

func ForwardToReplicas(v bool) *ForwardToReplicasOption {
	return &ForwardToReplicasOption{v}
}

func (o ForwardToReplicasOption) Get() bool {
	return o.value
}

func (o ForwardToReplicasOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *ForwardToReplicasOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = false
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *ForwardToReplicasOption) Equal(o2 *ForwardToReplicasOption) bool {
	if o2 == nil {
		return o.value == false
	}
	return o.value == o2.value
}

func ForwardToReplicasEqual(o1, o2 *ForwardToReplicasOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
