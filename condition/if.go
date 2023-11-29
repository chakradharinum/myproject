package condition

import "fmt"

func RunIfExample() {
	number := 10
	if number%2 == 0 {
		fmt.Println("Number is even.")
	} else {
		fmt.Println("Number is odd.")
	}
}
