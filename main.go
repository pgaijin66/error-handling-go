package main

import (
	"fmt"
	"log"
)

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

func errorReturningFucntion(flag bool) error {
	if !flag {
		fmt.Println("This function returns an error")
		return fmt.Errorf("flag is set to false")
	}
	return nil
}

func main() {
	flag := false
	start()
	fmt.Println("Application recoverd and regained contropl and application can still continue")

	err := errorReturningFucntion(flag)
	if err != nil {
		log.Fatal("Error occured in function errorReturningFunction", err)
	}
}
