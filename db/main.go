package db

import (
	"fmt"
	"net"
	"strings"

)

type Handler func(args []string) string 



var Commands = map[string]Handler{
    "set": Set,
    "get": Get,
    "del":Delete,
    "keys":Keys,
    "flushall":FlushAll,
    "hmset":HmSet,
    "hget":GetHashes,
    "hgetall":HgetAll,
    "hexists":Exists,
    "hdel":Hdel, 
    "setex":Set,
    "lpush":LPush,
    "ldisplay":DisplayList,
}



var DB = &Data{
    Prims: make(map[string]interface{}),
    Hashes: make(Hashes),
    List: make(List),
}


func Run(command string, conn net.Conn){


    command = strings.TrimSpace(strings.ToLower(command))

    tokens := strings.Split(command, " ")
    found := false 
    for cmd, fn := range Commands{
        if cmd == tokens[0]{
            err := fn(tokens[1:])
            fmt.Fprintf(conn, "%v\n", err)
            found = true
        }

        }       
        
        if !found{
            fmt.Fprint(conn, "Command not found")
        }

}



