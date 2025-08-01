package logicpack

import "fmt"

func PanicHandler() {
	msg := recover()

	fmt.Printf("Telah terjadi error : %s \n", msg)
}
