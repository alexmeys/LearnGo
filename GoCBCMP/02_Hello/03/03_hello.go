package main

func nope() {
	const ok = true
	var hello = "hello"
	_ = hello
}

func main() {
	// does not print, different scope/function.
	//fmt.Println(hello, ok)
}
