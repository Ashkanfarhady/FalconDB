package handlers

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

//TODO: randomized values for testcases

func createCommand(params ...string) string {
	return fmt.Sprintf("%s\r\n", strings.Join(params, "\r\n"))
}

func createCommandIO(params ...string) *bufio.Reader {
	content := createCommand(params...)
	stringReader := strings.NewReader(content)
	return bufio.NewReader(stringReader)
}

func createEmptyDB() map[string]string {
	return make(map[string]string)
}

func TestGetHandler(t *testing.T) {
	fdb := NewFalconDB()

	fdb.db["x"] = "data"
	command := createCommandIO("$1", "x")
	result := fdb.GetHandler(command)
	expected := createCommand("$4", "data")

	if result != expected {
		t.Fatalf("Expected '%s' but result is '%s'", result, expected)
	}

	command = createCommandIO("$1", "y")
	result = fdb.GetHandler(command)
	expected = createCommand("$-1")

	if result != expected {
		t.Fatalf("Expected '%s' but result is '%s'", result, expected)

	}
}

func TestSetHandler(t *testing.T) {
	fdb := NewFalconDB()

	command := createCommandIO("$1", "x", "$1", "1")
	result := fdb.SetHandler(command)
	expected := createCommand("+OK")

	if result != expected {
		t.Fatalf("Expected '%s' but result is '%s'", result, expected)

	}
	if fdb.db["x"] != "1" {
		t.Fatalf("Value db[\"x\"] was expecting to be '%s' but it is '%s'", "1", fdb.db["x"])
	}
}
func TestDeleteHandler(t *testing.T) {
	fdb := NewFalconDB()
	fdb.db["x"] = "data"
	command := createCommandIO("$1", "x")
	result := fdb.DeleteHandler(command)
	expected := createCommand("+OK")

	if result != expected {
		t.Fatalf("Expected '%s' but result is '%s'", result, expected)

	}
	if _, exists := fdb.db["x"]; exists {
		t.Fatalf("Deleleting fail due '%s' is still exists with value %s", "x", fdb.db["x"])

	}
	command = createCommandIO("$1", "random")
	result = fdb.DeleteHandler(command)
	expected = createCommand("$-1")

	if result != expected {
		t.Fatalf("Expected '%s' but result is '%s'", result, expected)
	}
}

func TestCommandHandler(t *testing.T) {
	fdb := NewFalconDB()
	result := fdb.CommandHandler()
	if result != createCommand("") {
		t.Fail()
	}
}

func TestInterprationHandler(t *testing.T) {
	fdb := NewFalconDB()

	testCases := []struct {
		command string
		params  *bufio.Reader
		result  string
		db      map[string]string
	}{
		{
			"set",
			createCommandIO("$1", "x", "$3", "hey"),
			createCommand("+OK"),
			createEmptyDB(),
		},
		{
			"set",
			createCommandIO("$1", "x", "$3", "hey"),
			createCommand("+OK"),
			map[string]string{"x": "Hello!"},
		},
		{
			"get",
			createCommandIO("$1", "x"),
			createCommand("$-1"),
			createEmptyDB(),
		},
		{
			"get",
			createCommandIO("$1", "x"),
			createCommand("$6", "Hello!"),
			map[string]string{"x": "Hello!"},
		},
		{
			"del",
			createCommandIO("$1", "x"),
			createCommand("+OK"),
			map[string]string{"x": "Hello!"},
		},
		{
			"del",
			createCommandIO("$1", "x"),
			createCommand("$-1"),
			createEmptyDB(),
		},
	}

	for _, tc := range testCases {
		fdb.db = tc.db
		result := fdb.InterprationHandler(tc.command, tc.params)
		if result != tc.result {
			t.Fatalf("testcase %v failed!\r\nresult '%s' was expected but '%s' is received.", tc, tc.result, result)
		}
	}
}
