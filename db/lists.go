package db 

type List map[string][]interface{}

func (l List) Push(key string, elem ...any) {
    l[key] = make([]interface{}, len(elem))
    l[key] = append(l[key], elem...)
}


