package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "Lesha"
	userNames = append(userNames, "Tanya", "Chelsey")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9

	courseRatings.output()

	// for index, value := range userNames {
	// 	fmt.Println(index, value)
	// }

	for course, rating := range courseRatings {
		fmt.Println(course, rating)
	}

}
