package useful

func index[E comparable](array []E, item E) int {
	for i, s := range array {
		if s == item {
			return i
		}
	}
	return -1
}
