package accumulate

// Accumulate applies the given converter on each element of the input 
// slice and output the slice with converted element
func Accumulate(input []string, converter func(string) string) []string {
	var res []string
	for _, elem := range input {
		res = append(res, converter(elem))
	}

	return res
}
