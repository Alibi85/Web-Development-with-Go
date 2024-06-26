// package main
//
// import (
// 	"os"
//
// 	"github.com/go-mail/mail/v2"
// )
//
// // All of these values will vary depending on your mail service.
// const (
// 	host     = "sandbox.smtp.mailtrap.io"
// 	port     = 587
// 	username = "fcc4af6eebd549"
// 	password = "03ab06406cbc53"
// )
//
// func main() {
// 	from := "test@lenslocked.com"
// 	to := "jonah.purinton@gmail.com"
// 	subject := "This is a test email"
// 	plaintext := "This is the body of the email"
// 	html := `<h1>Hello there buddy!</h1><p>This is the email</p><p>Hope you enjoy it</p>`
//
// 	msg := mail.NewMessage()
// 	msg.SetHeader("To", to)
// 	msg.SetHeader("From", from)
// 	msg.SetHeader("Subject", subject)
// 	msg.SetBody("text/plain", plaintext)
// 	msg.AddAlternative("text/html", html)
// 	msg.WriteTo(os.Stdout)
//
// 	dialer := mail.NewDialer(host, port, username, password)
// 	err := dialer.DialAndSend(msg)
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	stdctx "context"
	"fmt"

	"github.com/joncalhoun/lenslocked/context"
	"github.com/joncalhoun/lenslocked/models"
)

func main() {
	ctx := stdctx.Background()

	user := models.User{
		Email: "jon@calhoun.io",
	}

	ctx = context.WithUser(ctx, &user)

	retrievedUser := context.User(ctx)
	fmt.Println(retrievedUser.Email)
}
