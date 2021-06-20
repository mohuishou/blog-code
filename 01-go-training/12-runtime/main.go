package main

func main() {
	data := a()
	_ = data
}

//go:noinline
func a() int {
	ta := 2
	c := &ta
	_ = c
	return ta
}

//go:noinline
func b() *int {
	tb := 2
	return &tb
}
