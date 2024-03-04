package main

import (
	"fmt"
	"neversitup-test/question1"
	"neversitup-test/question2"
	"neversitup-test/question3"
)

func main() {

	strToPermute := "abcd"
	answer1 := question1.Permutation(strToPermute)
	fmt.Printf("Question 1: Permutation. %#v \n", strToPermute)
	fmt.Printf("Answer:\t %v \n\n", answer1)

	intArr := []int{1, 1}
	answer2 := question2.CountOddOccur(intArr)
	fmt.Printf("Question 2: Find the odd int. %#v \n", intArr)
	fmt.Printf("Answer:\t %v \n\n", answer2)

	smileys := []string{":)", ";(", ";}", ":-D"}
	answer3 := question3.CountSmileys(smileys)
	fmt.Printf("Question 3: Count the smiley faces. %#v \n", smileys)
	fmt.Printf("Answer:\t %v \n\n", answer3)
}
