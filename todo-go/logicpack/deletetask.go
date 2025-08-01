package logicpack

import (
	"fmt"
	"todo-go/userdata"
)

func DeleteTask() error {
	ShowTask()

	fmt.Print("Pilih yang ingin dihapus : ")
	var choice int
	_, err := fmt.Scan(&choice)

	if err != nil {
		return err
	}

	if choice > len(userdata.Todos) {
		return &notFoundErr{Message: "data tidak ditemukan"}
	}

	newData := append(userdata.Todos[:choice], userdata.Todos[choice+1:]...)

	saveData := &userdata.Todos

	*saveData = newData
	return nil
}
