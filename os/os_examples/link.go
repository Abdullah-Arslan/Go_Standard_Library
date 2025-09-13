package main
import (
    "fmt"
    "os"
)
func main() {
    err := os.Link("demo.txt", "hardlink.txt")
    if err != nil {
        fmt.Println("Hata:", err)
    }
}