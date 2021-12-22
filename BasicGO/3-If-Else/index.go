package main

import "fmt"

func main() {
	
	someValue := 9
	if someValue == 10 {
		fmt.Println("==10")
	} else {
		fmt.Println("!=10")
	}

	if someValue == 10 || someValue < 20 {
		fmt.Println("someValue == 10 || someValue < 20")
	} else {
		fmt.Println("nope1")
	}

	if someValue >= 10 && someValue <= 20 {
		fmt.Println("someValue >= 10 && someValue <= 20")
	} else {
		fmt.Println("nope2")
	}

	if result := doSomething(); result == "ok" {
		fmt.Println(result)
	} else {
		fmt.Println("okn't")
	}

	switchCase()

}

func doSomething() string {
	//TODO
	return "ok"
}

func switchCase() {
	index := 5
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("something else")
		break
	}
}
