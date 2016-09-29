package s3store

import (
	"errors"
)

func newMultiError(errs []error) error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	}

	message := "Multiple errors occurred:\n"
	for _, err := range errs {
		message += "\t" + err.Error() + "\n"
	}
	return errors.New(message)
}
