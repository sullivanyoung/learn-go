package loops

import "fmt"

// Loops function
func Loops() {
	// x := 0

	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	stringTest := []string{"test1", "test2", "test3", "test4"}
	
	// for i := 0; i < len(stringTest); i++ {
	// 	fmt.Println(stringTest[i])
	// }

	// for index, value := range stringTest {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// } 

	for _, value := range stringTest {
		fmt.Printf("the value is %v \n", value)
		value = "test5"
	} 

	fmt.Println(stringTest)
}
