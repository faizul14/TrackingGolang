package main

import "fmt"

func main() {

	nama := "faezol padli"

	fmt.Printf("Name:\t %v\n", nama)

	var nama2 *string = &nama
	
	*nama2 = "faezolmp"

	fmt.Printf("Before Name:\t%v\nAfter Name:\t%v", nama, *nama2)
}

