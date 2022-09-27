package param

func nonEmptyString(item string) bool {
	return len(item) > 0
}

func nonEmptyKeyString2(item [2]string) bool {
	return len(item[0]) > 0
}

func nonEmptyKeyString3(item [3]string) bool {
	return len(item[0]) > 0
}
