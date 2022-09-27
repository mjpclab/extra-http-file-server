package util

func Filter[T any](inputs []T, filterFunc func(T) bool) (outputs []T) {
	outputs = make([]T, 0, len(inputs))
	for i := range inputs {
		if filterFunc(inputs[i]) {
			outputs = append(outputs, inputs[i])
		}
	}
	return
}

func Concat[T any](inputs ...[]T) (outputs []T) {
	allLen := 0
	for i, length := 0, len(inputs); i < length; i++ {
		allLen += len(inputs[i])
	}

	outputs = make([]T, 0, allLen)
	for i := range inputs {
		outputs = append(outputs, inputs[i]...)
	}

	return
}
