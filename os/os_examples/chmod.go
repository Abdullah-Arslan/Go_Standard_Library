package main
import (
    "fmt"
    "os"
)
func main() {
    f, _ := os.Create("demo.txt")
    f.Close()
    err := os.Chmod("demo.txt", 0644)
    if err != nil {
        fmt.Println("Hata:", err)
    } else {
        fmt.Println("İzinler değiştirildi")
    }
}