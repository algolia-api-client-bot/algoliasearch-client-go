// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractAttributeForDistinct returns the first found AttributeForDistinctOption from the
// given variadic arguments or nil otherwise.
func ExtractAttributeForDistinct(opts ...interface{}) *opt.AttributeForDistinctOption {
	for _, o := range opts {
		if v, ok := o.(*opt.AttributeForDistinctOption); ok {
			return v
		}
	}
	return nil
}
