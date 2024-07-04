package main

import (
	"bufio"
	"net"

	"github.com/myselfBZ/mydb/db"
	"github.com/myselfBZ/mydb/utils"
)



func main(){
    lis, err := net.Listen("tcp", ":8080")
    utils.FailOnErr(err, "error starting a server")
    for {
        conn, err := lis.Accept()
        utils.LogErr(err, "error connecting to the client")
        data,  err := bufio.NewReader(conn).ReadString('\n')
        utils.LogErr(err, "error reading data")
        db.Run(data, conn)        
    }
}
