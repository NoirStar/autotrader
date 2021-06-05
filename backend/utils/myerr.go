package utils

import (
	"fmt"

	"github.com/pkg/errors"
)

// CheckErr checking err
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("ERROR: %+v\r\n", errors.Wrap(err, ""))
	}
}
