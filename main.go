package main

import "fmt"

type DNSHeader struct {
	id          int
	flags       int
	questions   int
	answers     int
	authorities int
	additionals int
}

type DNSQuestion struct {
	name  byte
	type_ int
	class int
}

func main() {

	fmt.Println("hello world")

}
