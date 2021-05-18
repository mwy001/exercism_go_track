package secret

// Handshake checks the input int and returns the required string array
// based on the mapping algorithm 
func Handshake(input uint) []string {
	var res []string

	if input & 0x1 != 0 {
		res = append(res, "wink")
	}

	if input & 0x2 != 0 {
		res = append(res, "double blink")
	}

	if input & 0x4 != 0 {
		res = append(res, "close your eyes")
	}

	if input & 0x8 != 0 {
		res = append(res, "jump")
	}

	if input & 0x10 != 0 {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}

	return res
}
