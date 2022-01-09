package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	//fmt.Println(demomod.HelloString())

	//Print different data types
	fmt.Println("########Printing different data type variables")
	diff_datatypes()
	fmt.Println("#####Arthimetic operations#######")
	operations()

	fmt.Println("#####Pointer Function#######")
	var x int = 5
	pointer_example(&x)

	fmt.Println("pointer example: vaule of x ", x)

	for i := 1; i < 10; i++ {

		fmt.Println("Value of i is", i, " time of the day is", time.Now().Hour())

	}

	switch hours := time.Now().Hour(); {
	case hours == 18:
		fmt.Println("Play time")
		fallthrough
	case hours == 22:
		fmt.Println("Sleep time")
	case hours == 10:
		fmt.Println("Office time")
	case hours == 7:
		fmt.Println("Wake up time")
	default:
		fmt.Println("Relax & enjoy !!")
	}

	fmt.Println("#######Playing with arrays#######")

	var arrayA [3]int
	arrayA[0] = 101
	arrayA[1] = 102
	arrayA[2] = 103

	arrayB := []int{201, 202, 203}

	fmt.Println("Values of arrays", arrayA, arrayB)

	for i, v := range arrayA {
		fmt.Println("Values and index is ", i, " ", v)
	}

	for i, v := range arrayB {
		fmt.Println("Values and index is ", i, " ", v)
	}

	fmt.Println("#######Playing with slices#######")

	sliceA := make([]int, 3)
	sliceA[0] = 1001
	sliceA[1] = 1002
	sliceA[2] = 1003
	sliceA = append(sliceA, 1004)
	fmt.Println(sliceA)

	fmt.Println("#######Slice Operator#######")

	s0 := []int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}

	s1 := s0[2:4]

	fmt.Println("value of s0", s0)
	fmt.Println("value of s1", s1)

	s1[0] = 22 * 2

	fmt.Println("After operation, value of s0", s0)
	fmt.Println("After operation, value of s1", s1)

	s1 = append(s1, 34)

	fmt.Println("After append, value of s0", s0)
	fmt.Println("After append, value of s1", s1)

	fmt.Println("#######maps example#######")
	maps_example()

}

func maps_example() {

	mapA := make(map[string]int)

	mapA["Arsenal"] = 6
	mapA["ManU"] = 5
	mapA["Chelsea"] = 1

	team, ok := mapA["ManC"]

	fmt.Println(team)
	fmt.Println(ok)

}

func pointer_example(x *int) {

	*x = 10

}

func operations() {

	x, y := 30, 5

	println("x+y = ", x+y)
	println("x-y = ", x-y)
	println("x*y = ", x*y)
	println("x/y = ", x/y)
	println("x mod y = ", x%y)

	day := true
	var night bool = false

	println("day||night = ", day || night)
	println("day&&night = ", day && night)
	println("!day = ", !day)

}
func diff_datatypes() {

	var a int = 5
	var b float32 = 2.34
	const pi float32 = 3.141
	var str string = "I am string"

	println(a)
	println(b)
	println(pi)
	println(str)

}

