package main

import "fmt"

func main() {
	primes:=[]int{2,3,5,7,11,13,17,19}

	m := make(map[int]int)
	

//loop for make all primes map to 0
for i:=0; i<=7; i++ {
	m[primes[i]]=0
}

fmt.Println(m)
}


//function calculates factorial of a number
func fact(num int) int{
	prod:=1
	for i := 2; i <= num; i++ {
		prod = prod * i
	} 
	return prod
}