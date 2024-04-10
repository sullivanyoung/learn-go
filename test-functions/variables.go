package variables

import "fmt"

// Test Variables function
func TestVariables() string {
    // strings
    var nameOne string = "Sullivan"
    var nameTwo = "Young"
    var nameThree string

    fmt.Println(nameOne, nameTwo, nameThree)

    nameOne = "Sully"
    nameThree = "GoLang Test"

    fmt.Println(nameOne, nameTwo, nameThree)

    nameFour := "Sull"
    fmt.Println(nameFour)

    // ints
    var ageOne int = 25
    var ageTwo = 30
    ageThree := 35

    fmt.Println(ageOne, ageTwo, ageThree)

    // bits & memory
    // https://pkg.go.dev/builtin#:~:text=floating%2Dpoint%20numbers.-,type%20int%20%C2%B6,-type%20int%20int
    var numOne int8 = 25
    var numTwo int8 = -128
    // uint cannot be assigned negative value
    var numThree uint = 25

    fmt.Println(numOne, numTwo, numThree)

    // floats
    // have to specify bit size (32 / 64) with float
    var scoreOne float32 = -1.5
    var scoreTwo float64 = -5677777.435645667
    scoreThree := 1.5

    fmt.Println(scoreOne, scoreTwo, scoreThree)

    return nameOne
}
