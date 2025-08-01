package logicpack

import (
	"fmt"
	"todo-go/userdata"
)

func ShowTask() {

	// data := &userdata.Todos

	if len(userdata.Todos) == 0 {
		fmt.Println("Kamu masih belum memiliki tugas :(")
		return
	}

	for _, t := range userdata.Todos {
		done := "✅"

		if !t.IsDone {
			done = "❌"
		}

		fmt.Println(t.TaskName, done)

	}

	fmt.Println("Total tugas yang ada adalah ", len(userdata.Todos))

}
