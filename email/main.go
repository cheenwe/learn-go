package main

import (
	"fmt"

	"./send_mail" //import package
)

func main() {
	mycontent := "test golang email"

	email := sendemail.NewEmail("cxhyun@126.com",
		"test golang email", mycontent)

	err := sendemail.SendEmail(email)

	fmt.Println(err)

}
