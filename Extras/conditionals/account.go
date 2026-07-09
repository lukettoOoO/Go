package main

import "fmt"

func main() {
	balance := 522

	if balance < 0 {
		fmt.Println("Balance is below 0, add funds now or you will be charged a penalty.")
	} else if balance == 0 {
		fmt.Println("Balance is equal to 0, add funds soon.")
	} else {
		fmt.Println("Your balance is 0 or above.")
	}
}
