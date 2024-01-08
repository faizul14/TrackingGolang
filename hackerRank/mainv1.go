package main

import "fmt"


func logika(n int, data []int) []int {
	var result = []int{}
	var dataTempt = []map[string]int{}

	for i := 1; i <= n; i++ {
		for j := 0; j < len(data) - 1; j++ {
			if i == data[j+1] {
				dataTempt = append(dataTempt, map[string]int{"index": j+1, "value": data[j+1]})
			}
		}
	}

	for i := 0; i < len(dataTempt) - 1; i++ {
		for j := 0; j < len(data) - 1; j++ {
			if dataTempt[i]["index"] == data[j] {
				result = append(result, j+1)
			}
		}
	}

	fmt.Println(dataTempt)
	return result
}

func main() {
	var data = []int{4,3,1,2,5}

	fmt.Println(logika(len(data), data))
}
