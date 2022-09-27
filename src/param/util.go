package param

import (
	baseParam "mjpclab.dev/ghfs/src/param"
)

func nonEmptyString(item string) bool {
	return len(item) > 0
}

func nonEmptyKeyString2(item [2]string) bool {
	return len(item[0]) > 0
}

func nonEmptyKeyString3(item [3]string) bool {
	return len(item[0]) > 0
}

func toString3s(inputs []string) (outputs [][3]string) {
	allKeyValues := baseParam.SplitAllKeyValues(inputs)
	outputs = make([][3]string, len(allKeyValues))
	for i := range allKeyValues {
		copy(outputs[i][:], allKeyValues[i])
	}
	return outputs
}
