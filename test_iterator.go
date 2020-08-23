package main

import "fmt"

const HEIGHT = 9
const WIDTH = 9

func main() {
    for i := 1; i <= HEIGHT; i++ {
        for j := 1; j <= WIDTH; j++ {
            fmt.Print(i * j)
        }
        fmt.Println("")
    }
}
