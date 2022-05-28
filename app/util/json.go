package util

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
)

func ColorPrintJSON(j interface{}, compact bool, c *color.Color) error {
	var out []byte
	var err error

	if !compact {
		out, err = json.MarshalIndent(j, "", "    ")
	} else {
		out, err = json.Marshal(j)
	}

	if err == nil {
		_, _ = c.Println(string(out))
	}

	return err
}

func FPrintJSON(w io.Writer, j interface{}, compact bool) error {
	var out []byte
	var err error

	if !compact {
		out, err = json.MarshalIndent(j, "", "    ")
	} else {
		out, err = json.Marshal(j)
	}

	if err != nil {
		return err
	}

	_, _ = fmt.Fprintln(w, string(out))

	return nil
}
