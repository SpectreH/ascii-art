package checks

import "ascii-art/errors"

func CheckFile(err error) {
	if err != nil {
		errors.PrintErrorMessage(1)
	}
}
