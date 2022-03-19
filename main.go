package main

import (
	// "fmt"
	// "log"
	// "fmt"
	"todo_app/app/models"
	// "todo_app/config"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test3"
	// u.Email = "test@example.com"
	// u.Password = "test"
	// u.CreateUser()

	// user, _ := models.GetUser(2)
	// user.CreateTodo("First Todo")

	// u.Name = "test"
	// u.Email = "test@text.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	// fmt.Println(todos)

	// user2, _ := models.GetUser(2)
	// todos, _ := user2.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}