package validator

func Validate(rpnValue string) bool {
	var count = 0
	for i := 0; i < len(rpnValue); i++ {
		if rpnValue[i] == '(' {
			count++
		} else if rpnValue[i] == ')' {
			if count < 0 {
				return false
			}
			count--
		}
	}

	if count == 0 {
		return true
	}
	return false
}
