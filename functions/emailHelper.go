package functions

import (
	"net/mail"
	"strings"
)

func isStringEmail(str string) bool {
	index:=strings.Index(str,"@")
	return index > -1
}

func validateEmail(email string) bool{
	_,err:=mail.ParseAddress(email)
	return err==nil

}