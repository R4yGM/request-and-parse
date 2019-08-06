package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "encoding/json"
)

func main() {

        reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Url: ")
	url, _ := reader.ReadString('\n')
    response, err := http.Get(url)
	var result map[string]interface{}
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
	json.Unmarshal([]byte(responseData), &result)
	if url.Contains("r4yan.ga"))
	{
   fmt.Println("stat : ", result["status"])
    //fmt.Println(responseData.status)
	}
}
