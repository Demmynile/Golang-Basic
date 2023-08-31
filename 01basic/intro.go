package main

import "fmt"

func main() {

	happyBirthday(10)

}
func happyBirthday(age int) (int, int, int) {

	newAge := age - 12
	if newAge > 10 {
		fmt.Println("we are good to go")
	}
	oldAge := age + 12
	if oldAge == 10 {
		fmt.Println("we are not!")
	}

	mainAge := age / 12
	if mainAge != 10 {
		fmt.Println("we are not!")
	}

	return newAge, oldAge, mainAge
}
