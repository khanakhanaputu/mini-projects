package logic_app

import (
	"contact-go/storage"
	"errors"
	"fmt"
	"strconv"
)

func AddContact() error {
	var (
		newName   string
		newNumber string
	)
	fmt.Print("Masukkan Nama Contact: ")
	_, err := fmt.Scan(&newName)
	if err != nil || newName == "" {
		return err
	}
	fmt.Print("Masukkan Nomber: ")
	fmt.Scan(&newNumber)
	num, err := checkNumber(newNumber)
	if err != nil {
		return err
	}
	err = checkContactExist(newNumber, num)
	if err != nil {
		return err
	}
	newData := storage.Contact{Name: newName, Number: uint64(num)}
	sendData := &storage.Contacts
	*sendData = append(*sendData, newData)

	return nil
}

func checkContactExist(newName string, num uint64) error {
	for _, v := range storage.Contacts {
		if v.Name == newName || v.Number == uint64(num) {
			return errors.New("kontak sudah ada")
		}
	}
	return nil
}

func checkNumber(newNumber string) (uint64, error) {
	num, err := strconv.Atoi(newNumber)
	if err != nil || len(newNumber) < 8 || len(strconv.Itoa(num)) < 8 {
		return 0, errors.New("tolong masukkan nomor dengan benar, minimal 8 angka ya")
	}
	return uint64(num), nil
}
