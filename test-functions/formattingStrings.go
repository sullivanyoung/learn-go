package formattingStrings

import "fmt"

// Formatting strings function
func FormattedStrings() {
	age := 26
	name := "sullivan"

	// Print (does not automatically make new line)
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println (auto generates new line)
	fmt.Println("hello")
	fmt.Println("world")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (Formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you score %f points \n", 225.55)
	fmt.Printf("you score %0.1f points \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
