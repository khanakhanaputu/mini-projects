package main

import (
	"fmt"
	"todo-go/logicpack"
)

func main() {
	for {

		fmt.Println("Silahkan pilih menu yang ada ")
		fmt.Println("1. Tambah Tugas baru")
		fmt.Println("2. Lihat tugas yang ada")
		fmt.Println("3. Selesaikan tugas")
		fmt.Println("4. Hapus tugas")
		var choice int
		fmt.Print("Pilihan : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			err := logicpack.CreateTask()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			logicpack.ShowTask()
		case 3:
			err := logicpack.DoneTask()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			err := logicpack.DeleteTask()
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Input Yang dimasukkan salah!")
		}
	}
}
