package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type Outputable interface {
	saver
	Display()
}

func main() {
	result := add(1, 5)
	fmt.Println(result + 55)

	printSomething(12)
	printSomething(12.12)
	printSomething("string")
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todoText := getTodoData()
	userTodo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	// userNote.Display()
	err = outputData(userNote)

	if err != nil {
		return
	}

	// userTodo.Display()
	// err = saveData(userTodo)
	err = outputData(userTodo)

	if err != nil {
		return
	}
}

func printSomething(value interface{}) {
	typedValue, ok := value.(int)

	if ok {
		fmt.Println("Integer:", typedValue)
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func outputData(data Outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the data failed.")
		return err
	}

	fmt.Println("Saving the data succeeded!")
	return nil

}

func getNoteData() (string, string) {
	title := getUserInput("Title:")
	content := getUserInput("Content:")

	return title, content
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getUserInput(promt string) string {
	fmt.Printf("%v ", promt)
	// var value string
	// fmt.Scanln(&value)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func add[T int | float64 | string](a, b T) T {
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)

	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }

	// aStr, aIsStr := a.(string)
	// bStr, bIsStr := b.(string)

	// if aIsStr && bIsStr {
	// 	return aStr + " " + bStr
	// }

	return a + b
}
