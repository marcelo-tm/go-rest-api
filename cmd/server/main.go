package main

import "fmt"

// Run - is going to be responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting application")
	return nil
}

func main() {
	fmt.Println("Go rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
