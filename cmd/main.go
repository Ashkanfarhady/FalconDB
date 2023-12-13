package main

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/Ashkanfarhady/FalconDB/handlers"
	"github.com/Ashkanfarhady/FalconDB/utils"
)

const welcome_encoded = "%0A%0A%20%20______%20%20%20%20_%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20_____%20%20____%20%20%0A%20%7C%20%20____%7C%20%20%7C%20%7C%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%7C%20%20__%20%5C%7C%20%20_%20%5C%20%0A%20%7C%20%7C__%20__%20_%7C%20%7C%20___%20___%20%20_%20__%20%7C%20%7C%20%20%7C%20%7C%20%7C_)%20%7C%0A%20%7C%20%20__%2F%20_%60%20%7C%20%7C%2F%20__%2F%20_%20%5C%7C%20%27_%20%5C%7C%20%7C%20%20%7C%20%7C%20%20_%20%3C%20%0A%20%7C%20%7C%20%7C%20(_%7C%20%7C%20%7C%20(_%7C%20(_)%20%7C%20%7C%20%7C%20%7C%20%7C__%7C%20%7C%20%7C_)%20%7C%0A%20%7C_%7C%20%20%5C__%2C_%7C_%7C%5C___%5C___%2F%7C_%7C%20%7C_%7C_____%2F%7C____%2F%20%0A%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%0A----------------------------------------------"

func main() {
	falconDb := handlers.NewFalconDB()
	welcome, _ := url.QueryUnescape(welcome_encoded)
	fmt.Println(welcome)
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	utils.CheckError(err)
	defer l.Close()

	for {
		c, err := l.Accept()
		utils.CheckError(err)
		go handleConnection(c, falconDb)
	}
}

func handleConnection(c net.Conn, falconDb *handlers.FalconDB) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		_, err := utils.ReadInteger(reader)
		if err != nil {
			return
		}
		command := utils.ReadString(reader)
		answer := falconDb.InterprationHandler(command, reader)
		c.Write([]byte(answer))
	}
}
