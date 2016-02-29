//Example of array, slices, maps and ranges
package main

import (
	"fmt"
)

func main() {
	//standard variable description
	var even [5]complex64
	even[0] = 2
	even[1] = 4
	even[2] = 6
	even[3] = 4
	even[4] = 8
	odd := [5]complex64{1i, 3i, 5i, 7i, 9i} //implicit array declaration
	fmt.Println(even, even[3])
	fmt.Println(len(odd), odd)

	var sumproduct complex64
	for i := 0; i < 5; i++ {
		sumproduct += even[i] * odd[i]
		fmt.Println(sumproduct)
	}
	fmt.Println(sumproduct)

}
