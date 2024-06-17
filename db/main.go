package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/myselfBZ/mydb/utils"
)

type Handler func(args []string) error 



var Commands = map[string]Handler{
    "set": Set,
    "get": Get,
    "del":Delete,
    "keys":Keys,
    "flushall":FlushAll,
    "hmset":HmSet,
}



var DB Data


func Run(){

    DB = Data{
        Hashes: make(Hashes),
    }
    reader := bufio.NewReader(os.Stdin) 
    for{
         
        fmt.Print("> ")
        input, err := reader.ReadString('\n') 
        utils.FailOnErr(err, "error reading from a line")
        input = strings.TrimSpace(strings.ToLower(input))
        if input == ""{
            continue
        }
         
        tokens := strings.Split(input, " ")
        found := false 
        for cmd, fn := range Commands{
            if cmd == tokens[0]{
                err := fn(tokens[1:])
                utils.LogErr(err, "")
                found = true
            }

        }       
        
        if !found{
            fmt.Println("Command not found")
        }
    }
}
