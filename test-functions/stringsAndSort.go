package stringsAndSort

import (
	"fmt"
	"sort"
	"strings"
)

// Strings and Sort function
func StringsAndSort() {
	// Strings std package
	greeting := "hello world"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "goodbye"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// The original string remains unchanged
	fmt.Println("original string value", greeting)

	// Sort std package
	ages := []int{25, 27, 30, 40, 15, 39}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"test2", "test3", "test31", "test04"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "test31"))
}
