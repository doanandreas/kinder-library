package validator

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func Unique[T comparable](values []T) bool {
	uniqueValues := make(map[T]bool)

	for _, v := range values {
		uniqueValues[v] = true
	}

	return len(uniqueValues) == len(values)
}

func ContainsEmptyString(values []string) bool {
	for _, v := range values {
		if v == "" {
			return true
		}
	}
	return false
}
