package db

import (
	"errors"
	"sync"
	"time"

	"github.com/myselfBZ/mydb/utils"
)

var mu sync.Mutex

//Basically, this the head of the database where everything is stored. I know naming conventions are just elegant

type Data struct{
    Prims map[string]interface{}
    Hashes Hashes
    List List 
}

//all the methods for CRUD operations 

func (d *Data) Set(key string, v interface{}, expiry int) {
    mu.Lock()
    d.Prims[key] = v 
    mu.Unlock()
    if expiry != 0{
        go func(){
            time.Sleep(time.Duration(expiry) * time.Second)
            mu.Lock()
            delete(d.Prims, key)
            mu.Unlock()
        }()

    }
}


func (d *Data) Get(key string) (interface{}, error) {
	v, ok := d.Prims[key]
	if !ok {
		return nil, errors.New("key not found")
	}

    return v, nil 
}

func (d *Data) Delete(keys []string) error{
    for _, key := range keys{
        if _, ok := d.Prims[key]; !ok{
            return utils.KeyNotFound
        }
        delete(d.Prims, key)
    }
    return nil 
}










