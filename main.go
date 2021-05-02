package main

import (
	"os"
	"strconv"
)

func convert(nums []int){
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
	args := os.Args[1:]
	argsArray := make([]int, len(args))
	for i := 0; i < len(args); i++{
		num, err := strconv.Atoi(args[i])
		if err == nil{
			argsArray[i] = num
		}
	}
	convert(argsArray)
}
