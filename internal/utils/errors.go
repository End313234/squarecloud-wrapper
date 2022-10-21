package utils

import (
	"errors"
	"fmt"
)

func ThrowSquareCloudAPIError(details string) error {
	return errors.New(fmt.Sprintf("SquareCloud: %s", details))
}
