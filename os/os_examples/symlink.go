package main
import (
    "fmt"
    "os"
)
func main() {
    err := os.Symlink("demo.txt", "link.txt")
    if err != nil {
        fmt.Println("Hata:", err)
        return
    }
    target, _ := os.Readlink("link.txt")
    fmt.Println("Link hedefi:", target)
}