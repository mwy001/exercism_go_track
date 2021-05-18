package sublist

type Relation string

func issublist(l1 []int, l2 []int) bool {
	len1 := len(l1)
	len2 := len(l2)

	if len1 == 0 {
		return true
	}

	if len2 == 0 {
		return false
	}

	if len1 > len2 {
		return false
	}

	for i := 0; i <= len2 - len1; i++ {
		var j = 0
		for ; j < len1; j++ {
			if l1[j] != l2[i + j] {
				goto Exit
			}
		}
        Exit:
		if j == len1 {
			return true
		}
	}
	return false
}

func Sublist(l1 []int, l2 []int) Relation {

	l1IsSublistOfl2 := issublist(l1, l2)
	l2IsSublistOfl1 := issublist(l2, l1)

	if l1IsSublistOfl2 && l2IsSublistOfl1 {
		return "equal"
	} else if l1IsSublistOfl2 {
		return "sublist"
	} else if l2IsSublistOfl1 {
		return "superlist"
	} else {
		return "unequal"
	}
}
