package main

import (
	"strings"
	"bufio"
    "fmt"
    "os"
)

var indexHashMap map[string]int = make(map[string]int)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
    operation := os.Args[1]
    key := os.Args[2]
    value := os.Args[3]

    switch operation {
        case "set":
            set(key, value)
        case "get":
            fmt.Printf(get(key))
        default:
            fmt.Printf("%s. is not a handle operation\n", operation)
    }


}

func set(key string, value string) {
    fmt.Println(key, value);
    //message := []byte()
    f, err := os.OpenFile("database.mars", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    check(err)
    fi, err := f.Stat()
    check(err)
    defer f.Close()

    _, errWrite := f.WriteString(fmt.Sprintf("%s, %s\n", key, value))
    indexKey(key,  (int(fi.Size())))
    check(errWrite)
}

func get(key string) (string){

    file, err := os.Open("database.mars")
    check(err)
    defer file.Close()
    
    offSet := indexHashMap[key]
    file.Seek(int64(offSet), 0)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var text = scanner.Text()
        if strings.Contains(text, key) {
            return strings.Split(text, ",")[1]
        }
    }

    if err := scanner.Err(); err != nil {
        check(err)
    }
    return ""
}

func indexKey(key string, offSet int) {
    indexHashMap[key] = offSet
}