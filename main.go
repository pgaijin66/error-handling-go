package main

import "fmt"

func thirdPartyFunction() {
	fmt.Println("error occured here and we need tp panic")
	panic("error occured")

}

func start() {
	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
			fmt.Println("Handling error: ", r)
		}
	}()

	thirdPartyFunction()
}

func main() {
	start()
	fmt.Println("Application recoverd and regained contropl and application can still continue")
}
