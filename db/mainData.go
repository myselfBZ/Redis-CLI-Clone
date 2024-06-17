package db

import (
	"errors"

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










