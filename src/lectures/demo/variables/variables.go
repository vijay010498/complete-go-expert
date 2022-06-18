package main

import "fmt"

func main() {
	var myName = "vijay"
	fmt.Println("My name is", myName, myName)

	// Type annotation
	var name string = "Pablo"
	fmt.Println("name =", name)

	//Create and assign shorthand
	userName := "admin"
	fmt.Println("username = ", userName)

	// un initialized variables
	var sum int
	fmt.Println("The sum is", sum)

	// compound assignment
	part1, other := 1, 5
	fmt.Println("part 1 is ", part1, "other is", other)

	//"idiom"
	part2, other := 2, 0
	fmt.Println("part 2 is", part2, "other is", other)

	// re-assign
	sum = part1 + part2
	fmt.Println("Sum =", sum)

	// Block assignment
	var (
		lessonName = "variables"
		lessonType = "Demo"
	)
	fmt.Println("Lesson Name=", lessonName)
	fmt.Println("Lesson Type=", lessonType)

	// compound assignment and ignore one variable
	word1, word2, _ := "hello", "world", "!" //"!" will be ignored
	fmt.Println(word1, word2)

}
