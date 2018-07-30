package main

import (
    "bufio"
    "fmt"
    "os"
//"io/ioutil"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
 f, err := os.OpenFile("/tmp/input.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
if err != nil {
    panic(err)
}

defer f.Close()

if _, err = f.WriteString(text); err != nil {
    panic(err)
}
}
