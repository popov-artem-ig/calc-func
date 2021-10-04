package structures

func Push(s string, stack []string) []string {
	return append(stack, s)
}

func Pop(stack []string) (string, []string) {
	if len(stack) > 0 {
		return stack[len(stack)-1], stack[:len(stack)-1]
	}
	return "", stack
}

func PopTwo(stack []string) (string, string, []string) {
	var stackOp []string
	var a string
	var b string
	b, stackOp = Pop(stack)
	a, stackOp = Pop(stackOp)
	if a == "" {
		a = "0"
	}
	return a, b, stackOp
}

func Peek(stack []string) string {
	if len(stack) == 0 {
		return ""
	}
	return stack[len(stack)-1]
}

func Count(stack []string) int {
	return len(stack)
}