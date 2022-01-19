package main

import (
	"fmt"
	"strconv"
	"encoding/json"
)

func main() {

	tmp := make(map[int]map[string]string)
	
	for i:=0; i<10; i++ {
	
		tmp[i] = map[string]string{ 
			"message" : "from name " + strconv.Itoa(i),
		}
	}
	
	result,err := json.Marshal(tmp)
	
	if err != nil {
		
		fmt.Println(err.Error())
		
	} else {
	
		fmt.Println(string(result))
	}
}