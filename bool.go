package valider

type Bool struct {
	value  bool
	field  string
	errors *map[string][]Error
}

func (v *Validator) Bool(value bool, field string) *Bool {
	return &Bool{value, field, v.Errors}
}

func (b *Bool) Equal(v bool) *Bool {
	if b.value != v {
		(*b.errors)[b.field] = append((*b.errors)[b.field], Error{ErrNotEqual, CodeNotEqual})
	}
	return b
}
