package main


import "fmt"

func main(){
	obj := map[string]int{"a":1, "b":2}
	newObj := make(map[string]int)
	for k,v := range obj {
		newObj[k] = v
	}
	newObj["b"] = 3
	fmt.Println(newObj)
}