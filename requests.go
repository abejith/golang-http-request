package main
import (
    "net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

type Message struct {
	userId int64
	id int64
	title string
	completed bool
}

func main() {
	//woeid := "2295424" //woe id for Chennai
    // rs, err := http.Get("https://www.metaweather.com/api/location/"+woeid)
    rs, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        panic(err) 
    }
    defer rs.Body.Close()
 
    bodyBytes, err := ioutil.ReadAll(rs.Body)
    if err != nil {
        panic(err)
    }
 
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	var data Message


	json.Unmarshal([]byte(bodyString), data)
	fmt.Println(data)

	
}
