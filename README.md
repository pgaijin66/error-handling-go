## Know when to use panic and when not to.

`defer` runs in rever order, i.e the defer statement defined last will be executed first

when to use panic : For serious issue in a major part of the code is buggy and like on startup of when it needs to connect to important service to work, then its okay to panic


## Using recover if you want to recover from failure 

When not to `recover`: When working with third party code which you do not have control over and issues with this do not want to crash the entire application.


`recover` : recover listens to panics and regains cotrol of execution, keeping application up and letting the flow continue.

recovers must be defined at the first of function so when afunction ends, due to a normal execution or due to a panic, the defer will be executed and recover will
check if some panic happened.

This way it will handle error and not crash the application.


Note: `Recovers only work if the panic happens on the same routine as they were defined.`

```
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

```

In the above code, main called start function which again called thirdpartyfuction. Now a panic occured in thirdPartyFunction, when this happens, panic checks if there are any defer functions and exucutes them. Since there is not any it kills the function, and flow is retured to start, and when start panics, it checks any defer and executes them. There is one and hence it executes. In our defer we have implemented recover hence it recovers from panic by stopping panic from killsing the function. Flow is then passed back to main function which continuees and exits gracefully 


## Only os.Exit in main, all other function should return an error

`log.Fatal` or `panic` called OS.Exit. All other functions should should return and error and handled gracefully

```
func errorReturningFucntion(flag bool) error {
	if !flag {
		fmt.Println("This function returns an error")
		return fmt.Errorf("flag is set to false")
	}
	return nil
}

func main() {
	flag := false

	err := errorReturningFucntion(flag)
	if err != nil {
		log.Fatal("Error occured in function errorReturningFunction", err)
	}
}
```

which gives you good formatted output and graceful exit

```
2022/10/31 10:25:26 Error occured in function errorReturningFunctionflag is set to false
exit status 1
```