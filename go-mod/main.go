package main

import (
	"fmt"

	"github.com/learn-golang/go-mod/models"
	"github.com/novalagung/gubrak"
)

func main() {
	user := models.User{"u001", "Hendrik"}
	fmt.Println(user)
	fmt.Println(gubrak.RandomInt(10, 20))
}
