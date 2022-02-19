package main

import "fmt"

func main() {
	var counter int = 7
	fmt.Println("howdi partner", counter)
	var arrayDemo = [20]string{"Tom", "Jerry"} //fixed lenght ones are called arrays
	fmt.Println("array", arrayDemo)
	arrayDemo[3] = "harry"
	fmt.Println("array is ", arrayDemo)

	var sliceDemo []string
	sliceDemo = append(sliceDemo, "Aphrem")
	fmt.Println("sliceDemo after append", sliceDemo)
	for i := 0; i < len(arrayDemo); i++ { //for loop demo
		fmt.Println(arrayDemo[i])
	}

}
