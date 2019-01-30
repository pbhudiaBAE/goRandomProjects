package main

import ("fmt"
		"io/ioutil")

func main() {
	str1, err1 := ioutil.ReadFile("A.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(string(str1))

	str2, err2 := ioutil.ReadFile("B.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(str2))

	err1 = ioutil.WriteFile("A.txt", str2, 0644)
	if err1 != nil {
		panic(err1)
	}

	err2 = ioutil.WriteFile("B.txt", str1, 0644)
	if err2 != nil {
		panic(err2)
	}

}
