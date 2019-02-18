package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

const (
	validUUIDLen         = 36
	uuidWithoutDashesLen = 32
)

func main() {
	cmdArgs := os.Args
	if len(cmdArgs) < 2 {
		fmt.Println("Usage ...")
		os.Exit(1)
	}

	newID, err := fixDashes(cmdArgs[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(newID)
}

func fixDashes(id string) (string, error) {
	if !validate(id) {
		return "", fmt.Errorf("invalid input")
	}
	if strings.Contains(id, "-") {
		return removeDashes(id), nil
	}
	return addDashes(id)
}

func validate(id string) bool {
	idLen := len(id)
	if idLen == 32 {
		return true
	}
	if _, err := uuid.Parse(id); err != nil {
		return false
	}
	return true
}

func removeDashes(id string) string {
	return strings.Replace(id, "-", "", -1)
}

func addDashes(id string) (string, error) {
	var sb strings.Builder
	var err error

	for i, v := range id {
		if i == 8 || i == 12 || i == 16 || i == 20 {
			_, err = sb.WriteString("-")
			if err != nil {
				return "", errors.New("unparsable id")
			}
		}
		_, err = sb.WriteRune(v)
		if err != nil {
			return "", errors.New("unparsable id")
		}
	}

	uuid := uuid.Must(uuid.Parse(sb.String()))
	return uuid.String(), nil
}
