/**Program that takes float or many floats and converts them to real positive floats
floatCheck takes singular float and converts to real float (positive)
sliceCheck takes an array of float64 making a new slice and converts negatives to positives and copies to slice
i.e transferrring all positives to new slice
**/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//floatCheck takes a single float and returns real float(always positive float)
func floatCheck(p float64) float64{
	if p<0{
		return -p
	} else{
		return p
	}
}

//randFloats takes a min and max float value and returns n random floats in a slice of float64elements
func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64() * (max - min) //math function for random flat
	}
	return res
}

//sliceCheck takes a slice of float64 elements and then returns them while adding to the list the real float64
// For example -64.1 is added to the list followed by 64.1 (real value) whereas positives are simply returned
func sliceCheck(b[] float64) []float64{
	q := make([]float64,len(b),cap(b))
	j:=0
	for _,x:= range b{
		if x<0{
			q=append(append(q,x),-x)
			j++
		}
		if x>=0{
			q=append(q,x)
			j++
		}
	}
	q=q[len(b):]//ignore the 0s/ will be the length of the original slice
	return q
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a:=randFloats(-200.10, 101.98, 5)
	fmt.Println(floatCheck(-20.123123))
	reg := sliceCheck(a)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reg)), ", "), ""))//Sprint is used to read float as string and using , as separator
	// If the [] brackets are to be removed replace empty brackets with "[]"


}
