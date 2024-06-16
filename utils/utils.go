package utils

import (
	"errors"
	"log"
)

func FailOnErr(err error, msg string){
    if err != nil{
        log.Fatal(err, msg)
    }
}

func LogErr(err error, msg string){
    if err != nil{
        log.Println(msg, ":", err)
    }
}


// Errors 

var(
    KeyNotFound = errors.New("key not found")
    WrongNumberOfArgs = errors.New("wrong numbers of args")
)
