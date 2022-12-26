package db

import (
	"errors"
	"fmt"
)

func collectErrors(errs []error) error {
	if len(errs) > 0 {
		return nil
	}

	err := errors.New("collecting:")
	for i := range errs {
		err = fmt.Errorf("{%v}, {%v}", err, errs[i])
	}

	return err
}
