package main

/*
int add(int a, int b)
{
    return a + b;
}
*/
import "C"

func main() {
	s := C.add(1, 2)
	println("add(1,2) =>", s)
}
