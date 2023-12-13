package utils

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)


func ReadInteger(reader *bufio.Reader) (int, error) {
	commandCountString, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	commandCount, err := strconv.Atoi(strings.TrimSpace(commandCountString[1:]))
    if err != nil {
        return -1, err
    }
	return commandCount, err
}

func ReadString(reader *bufio.Reader) (string, error) {
	commandLength, _ := ReadInteger(reader)
	command := make([]byte, commandLength+2)
	_, err := io.ReadFull(reader, command)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(command)), nil
}
