package main

import "fmt"

func main() {

	var mapOne = map[string]string{
		"name": "faezol mp",
		"gender": "laki-laki",
		"birthday": "23",
	}

	fmt.Printf("Name: %v \n", mapOne["name"])
	sliceMap()

}

func sliceMap()  {
	var data = []map[string]string{
		{
			"name": "faezol padli",
			"bithday": "20",
		},
		{
			"name": "faezolmp",
			"bithday": "23",
		},
		{
			"name": "padli",
			"bithday": "20",
		},
	}

	data = append(data, map[string]string{"name":"fmp"})

	for _, v := range data {
		fmt.Println(v)
	}
}