package utilx

func If (condition bool, truev, falsev interface{}) interface{} {
	if condition {
		return truev
	} else {
		return falsev
	}
}
