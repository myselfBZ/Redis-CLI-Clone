package db

import (
	"errors"
	"fmt"

	"github.com/myselfBZ/mydb/utils"
)

type Data struct{
    Prims map[string]interface{}
    Hashes Hashes
}



func (d Data) Set(key string, v interface{}) {
    d.Prims[key] = v
}

func (d Data) Get(key string) (interface{}, error) {
	v, ok := d.Prims[key]
	if !ok {
		return nil, errors.New("key not found")
	}

    return v, nil 
}

func (d Data) Delete(keys []string) error{
    for _, key := range keys{
        if _, ok := d.Prims[key]; !ok{
            return utils.KeyNotFound
        }
        delete(d.Prims, key)
    }
    return nil 
}


type Hashes map[string]map[string]string


func (h Hashes) Set(keyValue []string) error {
    lenght := len(keyValue)
    isKey := true 
    var key string 
    if lenght % 2 != 1{
        return utils.WrongNumberOfArgs
    }
    h[keyValue[0]] = make(map[string]string)
    for i := 1; i < lenght; i++{
        if isKey{
            key = keyValue[i]
            isKey = false 
        } else {
            h[keyValue[0]][key] = keyValue[i]
            isKey = true 
        }
    }
    fmt.Println(h[keyValue[0]])
    return nil 
} 

func (h Hashes) Get(key, field string) (string, error){
    v, ok := h[key][field]
    if !ok{
        return "", utils.KeyNotFound
    }
    return v, nil 
}







