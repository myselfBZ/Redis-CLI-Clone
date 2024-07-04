package db

import(
    "github.com/myselfBZ/mydb/utils"
    "fmt"
)

// Hash - stores other key-value pairs with a unique key

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

func (h Hashes) Delete(key string, fields []string) error {
    m, ok := h[key]
    if !ok{
        return utils.KeyNotFound
    }
    for _, k := range fields{
        _, ok := m[k]
        if !ok{
            return utils.KeyNotFound
        }
        delete(m, k)
        
    }
    return nil 
}


