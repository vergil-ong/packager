package test

import (
	"encoding/json"
	"fmt"
)

func test1() {
	str := "[{\"age\":18,\"sex\":\"man\",\"username\":\"user1\"},{\"age\":29,\"sex\":\"female\",\"username\":\"user2\"}]";
	var m []map[string]interface{}
	json.Unmarshal([]byte(str),&m)
	fmt.Println(m)
}
