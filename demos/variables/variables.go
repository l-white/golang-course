package main

import "fmt"

func main() {
	var myName string = "Laura"
	fmt.Println("Hello, ", myName + "!")
	
	userName := "Admin"
	fmt.Println("userName = ", userName)

	var sum int
	fmt.Println("sum = ", sum)

	part1, other := 1, 5
	fmt.Println("part 1 is ",part1, "other is ",other)

	sum = part1 + 2
	fmt.Println(sum)

	var (
		lessonName = "Variables"
		lessonType = "Demo"
	)
	fmt.Println("lessonName =", lessonName, "lessonType =", lessonType)

	word1, word2, _ := "hello", "world", "!"
	fmt.Println(word1, word2)
}