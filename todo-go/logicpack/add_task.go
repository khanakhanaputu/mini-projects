package logicpack

import (
	"errors"
	"fmt"
	"todo-go/userdata"
)

func CreateTask() error {
	var newTask string

	fmt.Print("Masukkan tugas baru : ")
	fmt.Scan(&newTask)

	if newTask != "" {
		todos := &userdata.Todos
		todo := userdata.Task{TaskName: newTask}
		*todos = append(*todos, todo)
		return nil
	}
	return errors.New("data tidak boleh kosong")

}
