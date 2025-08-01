package main

import (
	"contact-go/logic_app"
	"fmt"
)

func main() {
	for {
		fmt.Println("Contact Kontak")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View Contact")
		fmt.Println("3. Delete Contact")
		fmt.Println("4. Find Contact")
		fmt.Print("Masukkan pilihan: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Yang bener milihnya cok")
		}
		switch choice {
		case 1:
			err := logic_app.AddContact()
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			err := logic_app.ViewContacts()
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := logic_app.DeleteContact()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			err := logic_app.FindContact()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
