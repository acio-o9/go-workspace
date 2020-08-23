package main

import "fmt"
import "math/rand"
import "time"
import "os"
import "bufio"
import "strconv"

func main() {
    fmt.Println("Welcome janken game!")

    for startMenu() {
        fmt.Println(selectSign(), getRandomSign())
    }

    fmt.Println("Goodbye!")
}

func startMenu() bool {
    fmt.Println("Input key. 1:start, other:exit")
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Scan()
    text := stdin.Text()
    return text == "1"
}

func selectSign() int {
    fmt.Println("Select your sign. 0:gu-, 1:choki, 2:pa-")
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Scan()
    text := stdin.Text()
    var result, _ = strconv.Atoi(text)
    return result
}

func getRandomSign() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2)
}
