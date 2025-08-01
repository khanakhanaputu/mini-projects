package logicpack

import (
	"fmt"
	"todo-go/userdata"
)

func DoneTask() error {
	fmt.Println("Task Mana Yang Ingin diselesaikan?")
	for i, t := range userdata.Todos {

		var done string

		done = "❌"
		if t.IsDone {
			done = "✅"
		}

		fmt.Printf("%d %s %v \n", i+1, t.TaskName, done)
	}
	fmt.Print("Task : ")
	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		return err
	}
	if choice > len(userdata.Todos) {
		return &notFoundErr{Message: "data tidak ditemukan"}
	}
	data := &userdata.Todos[choice-1].IsDone

	if *data {
		*data = false
		ShowTask()
		return nil
	}
	*data = true
	ShowTask()
	return nil
}
