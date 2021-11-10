package main

func main() {
	/*
		IDK why this is not working
		var n int = 0

		do {
			if n == 5 {
				break
			}
			n++
			println(n)
		} while(i < 10);
	*/
	for i := 10; i > 0; i-- {
		if i == 5 {
			break
		}
		println(i)
	}
	for i := 0; i < 10; i++ {
		println(i)
	}
}
