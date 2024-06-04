package main

// To make this function callable from JavaScript,
// we need to add the: "export add" comment above the function
//
//export add
func Add(a, b int) int {
	return a + b
}

//export subtract
func Subtract(a, b int) int {
	return a - b
}

//export multiply
func Multiply(a, b int) int {
	return a * b
}

// func main() {
// 	fmt.Println("Go WebAssembly Calculator")
// }
