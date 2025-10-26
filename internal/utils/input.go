package utils

import (
	"bufio"
	"strconv"
	"strings"
)

func InputString(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func InputInt(reader *bufio.Reader) (int, error) {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}
