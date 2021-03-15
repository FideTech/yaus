package utils

//StringContains checks if the slice of strings contains the passed value
func StringContains(values []string, val string) bool {
	for _, v := range values {
		if v == val {
			return true
		}
	}

	return false
}
