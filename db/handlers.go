package db

import (
	"fmt"

	"github.com/myselfBZ/mydb/utils"
)

func Set(args []string) error {
	if len(args) != 2 {
		return utils.WrongNumberOfArgs
 	}
    
    DB.Set(args[0], args[1])
    fmt.Println("OK")
    return nil 
}

func Delete(args []string) error{
    if len(args) == 0{
        return utils.WrongNumberOfArgs
    }
    err := DB.Delete(args)
    if err != nil{
        return err 
    }
    fmt.Println("OK")
    return nil 
}

func Get(args []string) error {
    if len(args) != 1{
        return utils.WrongNumberOfArgs
    }

    v, err := DB.Get(args[0])
    if err != nil{
        return err 
    }
    fmt.Println(v)
    return nil 
}

func Keys(args []string) error {
    counter := 0
    if len(DB) == 0{
        fmt.Println("DB is empty")
    }
    for v := range DB{
        counter++
        fmt.Println(counter, ")", v)
    }
    return nil 
}








