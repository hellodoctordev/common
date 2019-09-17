package utils

func Contains(array []interface{}, item interface{}) bool {
	contains := false

	for _, value := range array {
		if item == value {
			contains = true
			break
		}
	}

	return contains
}