package logging
 
import (
    "fmt"
    "log"
    "os"
    "sync"
)
 
var (
    logger *log.Logger
    mu     sync.Mutex
)
 
func init() {
    file, err := os.OpenFile("/tmp/log_cent.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        fmt.Printf("Error opening log file: %v\n", err)
        return
    }
    logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}
 
func Log(message string) {
    mu.Lock()
    defer mu.Unlock()
    if logger != nil {
        logger.Println(message)
    }
}
