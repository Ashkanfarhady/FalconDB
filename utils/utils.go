package utils

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInteger(reader *bufio.Reader) (int, error) {
	commandCountString, err := reader.ReadString('\n')
	if err != nil {
		return -1, err
	}
	commandCount, err := strconv.Atoi(strings.TrimSpace(commandCountString[1:]))
	CheckError(err)
	return commandCount, err
}

func ReadString(reader *bufio.Reader) string {
	commandLength, _ := ReadInteger(reader)
	command := make([]byte, commandLength+2)
	io.ReadFull(reader, command)

	return strings.TrimSpace(string(command))
}
