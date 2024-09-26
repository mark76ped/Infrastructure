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

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	title, content:= getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	err = outputData(userNote)
}


func printSomething(value interface{}) { //any value is allowed with interface{}

	typedVal, ok := value.(string)
	if !ok {
		fmt.Println("It does not work", typedVal)
		return
	}

	intVal, ok := value.(int)
	if !ok {
		fmt.Println("It does not work", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if !ok {
		fmt.Println("It does not work", floatVal)
		return
	}
	// switch value.(type) {
	// case int: 
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// default:
	// 	fmt.Println("Type must be an Integer, Float, or String.")
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	saveData(data)
	return nil
}

func saveData(data saver) error{
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}


func getUserInput(prompt string) (string) {
	fmt.Printf("%v ", prompt)
	reader :=bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}