package server

import (
	"fmt"
	"strconv"
)

func parseID(arg string) (int, error) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("invalid id: %s", arg)
	}
	return id, nil
}
