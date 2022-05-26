package main

import "fmt"

func main() {
	sum := 0
	nums := [5]int{1, 2, 3, 4, 5}
	for _, t := range nums {
		sum = sum + t
	}
	fmt.Println(sum)

}

//payments := []types.Payment{
//	{
//		ID:     1,
//		Amount: 7,
//	},
//	{
//		ID:     2,
//		Amount: 8,
//	},
//	{
//		ID:     3,
//		Amount: 0,
//	},
//}
//
//max := realize.Max(payments)
//fmt.Println(max.ID)
