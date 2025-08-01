package logic_app

import (
	"contact-go/storage"
	"errors"
	"fmt"
)

func DeleteContact() error {
	var (
		contactName string
	)
	fmt.Print("Masukkan nama yang ingin dihapus: ")
	_, err := fmt.Scan(&contactName)
	if err != nil {
		return err
	}
	for i, v := range storage.Contacts {
		if contactName == v.Name {
			storage.Contacts = append(storage.Contacts[:i], storage.Contacts[i+1:]...)
			fmt.Println("Kontak berhasil dihapus!")
			return nil
		}
	}
	return errors.New("kontak tidak ditemukan")
}
