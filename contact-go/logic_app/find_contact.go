package logic_app

import (
	"contact-go/storage"
	"errors"
	"fmt"
)

func FindContact() error {
	fmt.Print("Cari kontak : ")
	var choice string
	_, err := fmt.Scan(&choice)
	if err != nil {
		return err
	}

	for _, v := range storage.Contacts {
		if choice == v.Name {
			fmt.Println("======Kontak Ditemukan======")
			fmt.Printf("Nama 	: %s \n", v.Name)
			fmt.Printf("Nomer 	: %d \n", v.Number)
			fmt.Println("============================")
			return nil
		}
	}

	return errors.New("data tidak ditemukan")
}
