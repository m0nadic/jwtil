package util

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"io/ioutil"

	"regexp"
)

var boldGreen = color.New(color.FgGreen, color.Bold)
var boldYellow = color.New(color.FgYellow, color.Bold)

func PrintToken(r io.Reader, compact bool) error {

	tokData, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// trim possible whitespace from token
	tokData = regexp.MustCompile(`\s*$`).ReplaceAll(tokData, []byte{})

	token, err := jwt.Parse(string(tokData), nil)

	if token == nil {
		return fmt.Errorf("malformed token: %w", err)
	}

	// Print the token details
	color.Cyan("Header:")

	if err := ColorPrintJSON(token.Header, false, boldGreen); err != nil {
		return err
	}

	color.Cyan("Claims:")

	if err := ColorPrintJSON(token.Claims, false, boldYellow); err != nil {
		return err
	}

	return nil
}
