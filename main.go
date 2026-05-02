package main

type transformFn func(int) int

// func main() {
// 	numbers := []int{1, 2, 3, 4}
// 	doubled := transformNumbers(&numbers, getTransformerFunction())
// 	triple := transformNumbers(&numbers, func(number int) int { return number * 2 })
// 	quadruple := transformNumbers(&numbers, createTransformer(4))

// 	fmt.Println(doubled)
// 	fmt.Println(triple)
// 	fmt.Println(quadruple)

// }

// func transformNumbers(numbers *[]int, transform transformFn) []int {
// 	dNumbers := []int{}
// 	for _, value := range *numbers {
// 		dNumbers = append(dNumbers, transform(value))
// 	}

// 	return dNumbers
// }

// func double(numner int) int {
// 	return numner * 2
// }

// func tripl(numner int) int {
// 	return numner * 3
// }

// func getTransformerFunction() transformFn {
// 	return double
// }

// func createTransformer(factor int) func(int) int {
// 	return func(number int) int {
// 		return number * factor
// 	}
// }
