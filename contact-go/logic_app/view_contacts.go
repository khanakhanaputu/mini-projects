package logic_app

import (
	"contact-go/storage"
	"fmt"
)

func ViewContacts() error {
	if len(storage.Contacts) == 0 {
		fmt.Println("Kamu nolep, tambah kontak gih")
		return nil
	}
	for _, v := range storage.Contacts {
		fmt.Println("======================")
		fmt.Println("Nama : ", v.Name)
		fmt.Println("Nomer : +62", v.Number)
		fmt.Println("======================")
	}
	fmt.Printf("Kamu mempunyai %d kontak \n", len(storage.Contacts))
	return nil
}
