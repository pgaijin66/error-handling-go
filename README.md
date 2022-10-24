`defer` runs in rever order, i.e the defer statement defined last will be executed first

### When to use panic

For serious issue in a major part of the code is buggy and like on startup of when it needs to connect to important service to work, then its okay to panic


### When not to panic

When working with third party code which you do not have control over and issues with this do not want to crash the entire application.

### Using recover

`recover` : recover listens to panics and regains cotrol of execution, keeping application up and letting the flow continue.

recovers must be defined at the first of function so when afunction ends, due to a normal execution or due to a panic, the defer will be executed and recover will
check if some panic happened.

This way it will handle error and not crash the application.


### Note:

`Recovers only work if the panic happens on the same routine as they were defined.`

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