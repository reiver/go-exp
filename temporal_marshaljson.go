package tmp

import (
	"encoding/json"

	"github.com/reiver/go-erorr"
)

var _ json.Marshaler = Nothing[bool]()
var _ json.Marshaler = Nothing[string]()

// MarshalJSON makes it so json.Marshaler is implemented.
func (receiver Temporal[T]) MarshalJSON() ([]byte, error) {
	switch interface{}(receiver.value).(type) {
	case bool, string,json.Marshaler:
		// these are OK.
	default:
		return nil, erorr.Errorf("tmp: cannot marshal something of type %T into JSON because parameterized type is ‘%T’ rather than ‘bool’, ‘string’, or ‘json.Marshaler’", receiver, receiver.value)
	}

	if receiver.isnothing() {
		return nil, erorr.Errorf("tmp: cannot marshal opt.Nothing[%T]() into JSON", receiver.value)
	}

	if receiver.IsDefunct() {
		return nil, erorr.Errorf("tmp: cannot marshal defunct %#v into JSON", receiver.value)
	}

	return json.Marshal(receiver.value)
}
