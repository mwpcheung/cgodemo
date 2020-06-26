package main

/*
#include "uid.h"
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -luid -framework Foundation
*/
import "C"
import (
	"log"
	"unsafe"

	uuid "github.com/satori/go.uuid"
)

func pointer(data []byte) unsafe.Pointer {
	if data == nil || len(data) == 0 {
		return C.NULL
	}
	return unsafe.Pointer(&data[0])
}
func charpointer(data []byte) *C.char {
	if data == nil || len(data) == 0 {
		return (*C.char)(C.NULL)
	}
	return (*C.char)(unsafe.Pointer(&data[0]))
}

//NewUUID new uuid
func NewUUID() uuid.UUID {
	// return uuid.Must(uuid.NewV4())
	uidbytes := make([]byte, 16)
	C.NewUUID(charpointer(uidbytes))
	u, err := uuid.FromBytes(uidbytes)
	if err != nil {
		log.Printf("foundation create uid failed %v", err)
	}
	return u
}

func main() {
	log.Printf("generate uuid from foundation %s", NewUUID().String())
}
