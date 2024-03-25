package main

//
// import (
// 	"fmt"
// 	"html/template"
// 	"os"
// )
//
// type User struct {
// 	Name string
// 	// 	Age  int
// 	// 	Meta UserMeta
// 	Bio template.HTML
// }
//
// type UserMeta struct {
// 	Visits int
// }
//
// func main() {
// 	t, err := template.ParseFiles("hello.gohtml")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	user := User{
// 		Name: "Jon Calhoun",
// 		Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
// 	}
//
// 	err = t.Execute(os.Stdout, user)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
//
