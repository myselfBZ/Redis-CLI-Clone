package db

import (
	"errors"

	"github.com/myselfBZ/mydb/utils"
)

type Data map[string]interface{}

func (d Data) Set(key string, v interface{}) {
    d[key] = v
}

func (d Data) Get(key string) (interface{}, error) {
	v, ok := d[key]
	if !ok {
		return nil, errors.New("key not found")
	}

    return v, nil 
}

func (d Data) Delete(keys []string) error{
    for _, key := range keys{
        if _, ok := d[key]; !ok{
            return utils.KeyNotFound
        }
        delete(d, key)
    }
    return nil 
}







