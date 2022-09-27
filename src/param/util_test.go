package param

import "testing"

func TestToString3s(t *testing.T) {
	var inputs []string
	var outputs [][3]string

	inputs = []string{"#aa#bb#cc", ":dd:ee:ff"}
	outputs = toString3s(inputs)
	if len(outputs) != 2 {
		t.Error(len(outputs))
	}
	if outputs[0][0] != "aa" || outputs[0][1] != "bb" || outputs[0][2] != "cc" {
		t.Error(outputs[0])
	}
	if outputs[1][0] != "dd" || outputs[1][1] != "ee" || outputs[1][2] != "ff" {
		t.Error(outputs[1])
	}
}
