package db

import (
	"encoding/json"
	"strconv"

)


// Handlers for requests that are coming in from the outside network

// Elegant err messages
var wrongPattern = "Wrong pattern"


func Set(args []string) string {
    length := len(args)
    if length == 3{
        sec, err := strconv.Atoi(args[2])
        if err != nil{
            return "Invalid timeduration"
        }
        DB.Set(args[0], args[1], sec)
        return "" 
    } else if length == 2{
        DB.Set(args[0], args[1], 0)
        return "OK"
    }
    return "wrong patter"
}

func Delete(args []string) string{
    if len(args) == 0{
        return wrongPattern
    }
    err := DB.Delete(args)
    if err != nil{
        return err.Error() 
    }
    return "OK" 
}

func Get(args []string) string {
    if len(args) != 1{
        return wrongPattern
    }

    v, err := DB.Get(args[0])
    if err != nil{
        return err.Error() 
    }
    return v.(string) 
}

func Keys(args []string) string {
    if len(DB.Prims) == 0{
        return "db is empty"
    }
    data, _ := json.Marshal(DB.Prims)
    return string(data) 
}


func FlushAll(args []string) string {
    for k := range DB.Prims{
        delete(DB.Prims, k)
    }
    return "OK" 
}

func HmSet(args []string) string{
    err := DB.Hashes.Set(args)
    return err.Error() 
}

func GetHashes(args []string) string{
    if len(args) != 2{
        return wrongPattern
    }
    v ,err := DB.Hashes.Get(args[0], args[1])
    if err != nil{
        return err.Error()
    }
    return v
}


func HgetAll(args []string) string{
    m, ok := DB.Hashes[args[0]]
    if !ok{
        return "key not found" 
    }
    data, _ := json.Marshal(m)
    return string(data) 
}

func Exists(args []string) string {
    if len(args) != 2{
        return wrongPattern 
    }
    _, ok := DB.Hashes[args[0]][args[1]]
    if ok{
        return "true"
    }  
    return "false"
}


func Hdel(args []string) string{
    err := DB.Hashes.Delete(args[0], args[1:])
    return err.Error()
}


func DisplayList(args []string) string{
    if len(args) != 1{
        return wrongPattern
    }
    l, ok := DB.List[args[0]]
    if !ok{
        return "not found" 
    }
    data, _ := json.Marshal(l)
    return string(data)
}


func LPush(args []string) string{
    if len(args) > 1{
        DB.List.Push(args[0], args[1:])
        return "OK" 
    }
    return wrongPattern 
}








