package main
import (
    "os"
    "time"
)
func main() {
    now := time.Now()
    os.Chtimes("demo.txt", now, now)
}