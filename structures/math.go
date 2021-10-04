package structures

func GetOperation(operator string) func(x, y float64) float64 {
	switch operator {
	case "+":
		return Sum
	case "-":
		return Sub
	case "*":
		return Multi
	case "/":
		return Divide
	default:
		return nil
	}
}

func Sum(x, y float64) float64 {
	return x + y
}

func Sub(x, y float64) float64 {
	return x - y
}

func Multi(x, y float64) float64 {
	return x * y
}

func Divide(x, y float64) float64 {
	return x / y
}