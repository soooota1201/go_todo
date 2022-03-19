package main

import (
	"fmt"
	// "log"
	// "fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
	// "todo_app/config"
)

func main() {

	fmt.Println(models.Db)

	controllers.StartMainServer()
}