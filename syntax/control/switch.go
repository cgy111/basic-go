package main

func Switch(status int) {
	switch status {
	case 0:
		println("初始化")

	case 1:
		println("运行中")

	default:
		println("未知状态")

	}
}
func SwitchBool(age int) {
	switch {
	case age >= 18:
		println("成年人")

	case age > 12:
		println("青年")

		/*default:
		println("未知状态")
		*/
	}
	println("小孩子")
}
func SwitchAny(age int) {
	switch age {
	case 18:
		println("成年人")

	case 12:
		println("青年")

		/*default:
		println("未知状态")
		*/
	}
	println("小孩子")
}
