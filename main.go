package main

import (
	"strconv"
)

func convert(nums []int, length int){
	phonetic := map[string]string{
		"0" : "Zero",
		"1" : "One",
		"2" : "Two",
		"3" : "Three",
		"4" : "Four",
		"5" : "Five",
		"6" : "Six",
		"7" : "Seven",
		"8" : "Eight",
		"9" : "Nine",
	}
	for index := 0; index < len(nums); index++ {
		num := strconv.Itoa(nums[index])
		for _, c := range num{
			print(phonetic[string(c)])
		}
		if index == len(nums) - 1{
			break
		}
		print(", ")
	}
}
func main() {
	var sampleArray = []int{34, 65}
	convert(sampleArray, len(sampleArray))
}
