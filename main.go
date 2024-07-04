package main

import (
	"bufio"
	"log"
	"net"

	"github.com/myselfBZ/mydb/db"
	"github.com/myselfBZ/mydb/utils"
)



func main(){
    lis, err := net.Listen("tcp", ":8080")
    utils.FailOnErr(err, "error starting a server")
    log.Println("Server has started at localhost:8080")
    for {
        conn, err := lis.Accept()
        utils.LogErr(err, "error connecting to the client")
        data,  err := bufio.NewReader(conn).ReadString('\n')
        utils.LogErr(err, "error reading data")
        db.Run(data, conn)        
    }
}
