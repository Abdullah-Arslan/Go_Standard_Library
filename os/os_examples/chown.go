package main
import (
    "fmt"
    "os"
)
func main() {
    err := os.Chown("demo.txt", 1000, 1000)
    if err != nil {
        fmt.Println("Hata:", err)
    }
}