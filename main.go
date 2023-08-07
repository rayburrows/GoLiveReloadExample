package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	for {
		id := uuid.New()
		fmt.Println("########-> This is the new id: ", id.String())
		time.Sleep(4 * time.Second)
	}
}
