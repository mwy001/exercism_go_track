package flatten

func flatten(input []interface{}, res []interface{}) []interface{}{
	for _, i := range input {
		if i == nil {
			continue
		}
		switch v := i.(type) {
		case []interface{}:
			res = flatten(v, res)			
		default:
			res = append(res, i)
		}
	}

	return res
}

func Flatten(input interface{}) []interface{} {

	var res []interface{} = []interface{}{}

	switch v := input.(type) {
	case []interface{}:
		res = flatten(v, res)
	default:
	}
	return res
}
