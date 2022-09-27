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
