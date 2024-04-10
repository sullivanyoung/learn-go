package arraysAndSlices

import "fmt"

func ArraysAndSlices() {
	// Arrays (fixed length)
	var ages [3]int = [3]int{20, 25, 30}
	var agesTwo = [3]int{20, 25, 30}

	stringTest := [4]string{"test1", "test2", "test3", "test4"}
	stringTest[1] = "test2.0"

	fmt.Println(ages)
	fmt.Println(agesTwo, len(agesTwo))
	fmt.Println(stringTest, len(stringTest))

	// Slices (use arrays under the hood)
	var scores = []int{100, 50, 0}
	scores[2] = 25
	scores = append(scores, 0)

	fmt.Println(scores, len(scores))

	// Slice Ranges
	rangeOne := stringTest[1:3]
	rangeTwo := stringTest[2:]
	rangeThree := stringTest[:3]
	
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "test5")
	fmt.Println(rangeOne)
}
