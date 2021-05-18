package leap

// IsLeapYear checks whether the input year is leap year or not
func IsLeapYear(input int) bool {
	if input % 400 == 0 || (input % 4 == 0 && input % 100 != 0) {
		return true
	}
	return false
}
