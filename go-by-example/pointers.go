package main

import "fmt"

func change_arg_by_value(val int, newValue int) {
	val = newValue
}

func change_arg_by_reference(val *int, newVal int) {
	*val = newVal
}

func main() {
	var p *int // p reference to nil
	fmt.Println(p)

	i := 2
	p = &i          // p reference to addr i that hold 2
	fmt.Println(p)  // return address holding p
	fmt.Println(*p) // dereference p. return value p holding

	*p = 4          // re-assign value 4 to the currenr reference memory addr
	fmt.Println(p)  // still current memory address that hold i
	fmt.Println(*p) // dereference p.

	val := 1

	change_arg_by_value(val, 20) // original value of val will not change. passing by val
	fmt.Println(val)

	change_arg_by_reference(&val, 20) // original value of val is now changed. passing by reference
	fmt.Println(val)
}
