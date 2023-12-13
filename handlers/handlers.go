package handlers

import (
	"bufio"
	"fmt"
	"strings"
	"sync"

	"github.com/Ashkanfarhady/FalconDB/utils"
)

type FalconDB struct {
	db   map[string]string
	lock sync.RWMutex
}

func NewFalconDB() *FalconDB {
	falconDB := &FalconDB{}
	falconDB.db = make(map[string]string)
	return falconDB
}

func (f *FalconDB) GetHandler(reader *bufio.Reader) string {
	key, _ := utils.ReadString(reader)
	f.lock.RLock()
	defer f.lock.RUnlock()
	_, ok := f.db[key]
	if ok {
		return fmt.Sprintf("$%d\r\n%s\r\n", len(f.db[key]), f.db[key])
	} else {
		return "$-1\r\n"
	}
}

func (f *FalconDB) SetHandler(reader *bufio.Reader) string {
	key, _ := utils.ReadString(reader)
	value, _ := utils.ReadString(reader)
	f.lock.Lock()
	defer f.lock.Unlock()
	f.db[key] = value
	return "+OK\r\n"
}

func (f *FalconDB) DeleteHandler(reader *bufio.Reader) string {
	key, _ := utils.ReadString(reader)
	f.lock.Lock()
	defer f.lock.Unlock()
	_, ok := f.db[key]
	if ok {
		delete(f.db, key)
		return "+OK\r\n"
	} else {
		return "$-1\r\n"
	}

}

func (f *FalconDB) CommandHandler() string {
	return "\r\n"
}

func (f *FalconDB) InterprationHandler(command string, reader *bufio.Reader) string {
	command = strings.ToUpper(strings.TrimSpace(command))
	switch command {
	case "GET":
		return f.GetHandler(reader)
	case "SET":
		return f.SetHandler(reader)
	case "DEL":
		return f.DeleteHandler(reader)
	case "COMMAND":
		return f.CommandHandler()
	default:
		return "-Please provide a valid command.\r\n"
	}
}
