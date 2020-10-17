package main

import "fmt"

func Burbuja(s []int64) { //burbuja
	fmt.Println(s)
	for x := 0; x < len(s); x++ {
		for i, v := range s {
			if v != s[len(s)-1] {
				if s[i] > s[i+1] {
					temp := v
					s[i] = s[i+1]
					s[i+1] = temp
				}
			}
		}
	}
	fmt.Println(s)
}

func main() {
	slice := []int64{7, 4, 2, 8, 1, 5, 3, 6, 9}
	Burbuja(slice)
}
