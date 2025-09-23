package reflction

func Walk(x interface{}, fn func(input string)) {
	fn("Oh on")
}
