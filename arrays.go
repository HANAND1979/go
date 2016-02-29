//Example of array, slices, maps and ranges
package main

import (
	"fmt"
)

func main() {
	//Examples of range is provided for each of the data structures below
	
	//standard variable description
	var even [5]int
	even[0] = 2
	even[1] = 4
	even[2] = 6
	even[3] = 4
	even[4] = 8
	//even[5] = 10 //error will be thrown here. contrast this with how we can expand a slice dynamically later.
	odd := [5]int{1, 3, 5, 7, 9} //implicit array declaration
	fmt.Println(even, even[3])
	fmt.Println(len(odd), odd)

	var sumproduct int
	for i := 0; i < 5; i++ {
		sumproduct += (even[i] * odd[i])
		fmt.Println(sumproduct)
	}
	fmt.Println(sumproduct)
	
	//Example of use of range in the case of matrices 
	matrix:=make(map[int]int)
	for _,x:= range even {
		for _,y:=range odd{
			fmt.Println(x,y)		
			matrix[x]=x*y
			fmt.Println(matrix)
		}
	}
	fmt.Println(matrix) //notice from the output that there is no sorted order to a map

	//slices are parts of arrays
	s := make([]string, 5, 10) //declaring a slice
	fmt.Println(s)             //returns a blank matrix
	s[0] = "apple"
	s[1] = "bat"
	s[2] = "cat"
	fmt.Println(len(s)) //this returns the declared length of the slice 
	s = append(s, "dog")
	s = append(s, "eel")
	s = append(s, "fox")
	fmt.Println("set:", s)
	fmt.Println(len(s)) //observe how the length changes 
	s = append(s, "goat")
	s = append(s, "horse")
	s = append(s, "iguana")
	t:=make([]string,len(s))
	copy(t,s) //copying a slice 
	fmt.Println(len(s)) //observe how the length changes beyond size of 10 declared earlier 
	fmt.Println("get:", s[5])
	fmt.Println("appended set:", s)
	fmt.Println(s[2:5],s[2:],s[:9])
	fmt.Println("copied slice:",t)
	
	//Example of use of range for slices
	for q,r:= range s{
		fmt.Print(q,r)
	}
	
	//illustration of maps [key-value pairs]
	scrip:= make(map[string]float32)
	scrip["nifty"]=7000
	scrip["sensex"]=22000
	scrip["ITC"]=230
	fmt.Println(scrip)
	delete(scrip,"nifty")
	fmt.Println(scrip,len(scrip))
	_, is_present := scrip["nifty"] 
    	fmt.Println("is nifty present?:", is_present)
	
	//illustration of ranges
	for i,v:=range scrip{
		fmt.Println(i,v)
	}
	
}
