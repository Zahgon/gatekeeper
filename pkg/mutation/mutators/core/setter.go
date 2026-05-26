package core

// Setter tells the mutate function what to do once we have found the
// node that needs mutating.
type Setter interface {
	// SetValue takes the object that needs mutating and the key of the
	// field on that object that should be mutated. It is up to the
	// implementor to actually mutate the object.
	SetValue(obj map[string]interface{}, key string) error

	// KeyedListOkay returns whether this setter can handle keyed lists.
	// If it can't, an attempt to mutate a keyed-list-type field will
	// result in an error.
	KeyedListOkay() bool

	// KeyedListValue is the value that will be assigned to the
	// targeted keyed list entry. Unlike SetValue(), this does
	// not do mutation directly.
	KeyedListValue() (map[string]interface{}, error)
}

var _ Setter = &defaultSetter{}

func NewDefaultSetter(value interface{}) Setter { _ = "STUB: not implemented"; return *new(Setter) }

// defaultSetter is the default implementation of the Setter interface that supports
// assigning plain values and external data placeholders.
type defaultSetter struct {
	value interface{}
}

func (s *defaultSetter) KeyedListOkay() bool { _ = "STUB: not implemented"; return false }

func (s *defaultSetter) KeyedListValue() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *defaultSetter) SetValue(obj map[string]interface{}, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// make a copy of the incoming placeholder so we can modify it
