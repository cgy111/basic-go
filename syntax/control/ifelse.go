package control

func IfOnly(age int) {
	if age >= 18 {
		println("成年了")
	}
}
func IfElse(age int) string {
	if age >= 18 {
		println("成年了")
		return ""
	} else {
		println("小孩子")
	}
	return ""
}
func IfElseIf(age int) string {
	if age >= 18 {
		println("成年了")
		return ""
	} else if age >= 12 {
		println("青年")
	} else {
		println("小孩子")
	}
	return ""
}
