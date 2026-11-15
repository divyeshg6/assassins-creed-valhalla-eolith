package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.3" )

func LJDvPGibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6ezSr6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrIY4kX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R33Cjv8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zx5BPGdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZL3pnoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0uyrOOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrhCbuFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltkHXdZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGODx8PlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9thTp6AyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6N92iiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2igDin2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFQ3LwMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWYBBfBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXj3DY1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SwyBj9jzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nD0HdY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwQUHdnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2fxFxvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xADkAlLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHGj3blSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nI9h13nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZ0ihNMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orKuoDQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atwcKQA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGfhzdzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7fOmIzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSFGdIDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nm5z8MOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQmbO6SnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 716NIlQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOJS3XeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Dh3zAJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLohdYkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UE3XMqhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hE3xWdQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WITzGBkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPwuGaGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EzbehFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wj0jNywSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhG9YRfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gL5MJt3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmGBJn3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSPVrIo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayONpQ8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIDLGNg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9HZDXGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOxBeNKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igTs1sTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srgAProQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 09T76jyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQHJfGiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0ao3vLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0Un1HmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtK6puE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASGNZrMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfzbpP1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhIp9pT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFaHJk5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03fv05WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O17UrnraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szKP1brqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Xpoh3C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uQBJhxoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOa7Puk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjrUlo7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1axKCyM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUMky36AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6IRspraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiwjLr2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6AYtMRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxSBp8CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uZfD0KDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gimbhebuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUtkd90LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lf71BuR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeNALLeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QBpAgP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMfQYYPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JENnhUV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYg2ycJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2IBzjgRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oox6gd7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0oxmBRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Be5hK6ebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWODzA9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcbxKgjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3av6jNxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KFHRMrR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khQUstEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1djLE1X7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wH3HetB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjIX1EHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouPDapVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JR1DRZZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99oNSHCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGefUryUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IggQAOEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLp6q2kAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oK9L4T53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGlvP3qOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idryzG6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ztp0JXbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPlrQ5bbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMaMfgZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ot9ELSElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFcgqU1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HU3XW5asWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MnYa0VF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EImsJoMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiFOgINzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6gIZUarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axYDxiKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SntWdM55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUFLb95WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N157hfneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiZ31ddwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bl7yFDFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILtjssf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZunlpjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcGDYOD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXjYKfw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyMBUW8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrO7nHT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rp3ME9E5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTQdbCZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AG6TaSipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yawQlKjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AEJzwzevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IllZNwWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXugvzGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PXONldO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOV1pjH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzl1i5XeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMpWfOAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bw5r20YAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbUy5jxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rd3Gw6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITTpnLxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1EjlbAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KKA0sW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSiEMI2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrLohoPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JN7TDyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36hOU02eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEX1YYMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luBnaazZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func niGUHiHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HevHz0otWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJ9yLaRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHBbooAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rzfxDSeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wCJqAB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcDBzNxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sSwZWO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSe7FTR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIitaQScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lxi7BYN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rliZ0QjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5ceDTAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PraSfmbBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uN2iy9ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYJfonniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyY6denyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EQB0xfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkLtaNiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JyiL6nJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIIutimpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zaCY4BRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHVALHU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSmWgRm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQn8Vz00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func plIvycoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rADVpEPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4IniYnZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXXrH2QbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FV35TOJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tB2C5YzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktvFK4rfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfGHmc0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6f0idNhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXE2LJFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnKRF4eeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZiQhE6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PftpAb2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eq0Ku8nVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FfuoPMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGIGuIxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WZvYp0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bE8EsueNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSv5xQU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBnzQz3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vdWRJw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLSbd8kfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhyQiZA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lahpzZyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUKoNoULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIFt6tk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSWml1yjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KELrIqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nA4oO4tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRL3WJiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxnMLlciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4NMbWjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zr9Z9gAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEvajM7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mzxVRGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JT4N7yj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 492vDvsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZjGrTcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZDrw1HiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2v79PpvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I20eSFVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCGQPRN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZ6ODjpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSZfOrjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmixLohaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tawOXddKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNaOhOAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCkq6r9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ztsc8eLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuivgyMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2AnZjv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paaVlPSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Q6KwytiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfUapbXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFN8bnWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfXFDYCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krYon9y4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJM1XLhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8Y1Jh0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9fIGK6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDPMNZfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y18o8jKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iI8MyWNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sChpLJrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SO2N1OmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CV9pBff3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsBL3tzQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KW1C5ROFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Okud9wljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3Oj6ZC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUvdydj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2hgIjTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QfCEDR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lh8h7XloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNE7OReqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NakSBwgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ul7wJkdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDAnBh0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AeKEVnsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3c8MCwwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYJTEoVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woMxFidiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 575rZHDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4f8YcU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zR8jIdPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWvj5LcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X881myxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnweJtgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M47oeYLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfprI54gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IMQ1G7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ifOo7tLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxeJoAwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FuWpu77SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08kzrsb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjgkswryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZRleHUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDSNuA5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40q178RdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func De0Wu7KSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVqG6UcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVPwMtpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NztSkXOBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6FC2KltoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ottyEZP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKhfD9tRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCGh3UTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rlw4ujcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iloiBN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hg20Ub1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s3Ni2bK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aqk83mqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Pv3jFEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8aRqXTOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAy5CmugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lp4cJqOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1skwIwZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AYk7cEFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M09SIwkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ssWq5yJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tm3MsekZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0nUXrzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pf5YCF6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0R9cCyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0u01PU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CRpIdATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEdG6WHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjcsgkalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFS3rzuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96wBa5nIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PPJmenmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYE4Fhu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ry6Q1EKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSUx7BeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63ffdM0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9bOiAtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFd9kkUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XeLufkaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANAvMryNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JD6Dyj3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvgdXPl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1otwMHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FGu4b9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZgy3MRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ev5bnmnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zg1o1VTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGqXkOfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0g2UZgImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDJRYd2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p76iON1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dc6P3BGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdqLOpEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obBNI2rHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MqHesHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzcfOBsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjDRa9ShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGBe7Q47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gl93sZMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5RKojqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfZbrUEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwcuRJqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8TSLGBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1v5cdS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QlmMUwXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhuTkWLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2oNSXhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6e88ZfmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVSzLs01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIOiTLgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxRYTFH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuZo3HPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2KFwrY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKT3T1PLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmjAoroeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yb3TKPe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlbJrs12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nUzXNo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEBx5UZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQBBhqWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWFjL55hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qALxLFq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9b7OmuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taThkAmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FirfB0THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xnn0SOr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rk0jQThzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlAlQH6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idl9cwhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ga93jmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EW62pBtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSqgGDaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOI4JuC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDzdzdAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oj7BdEtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYOWFgv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2U3SCMwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6J9b03brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JP1bmvwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcbxNA8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLkKrzK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h40ZiGi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOI4AuFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFtFw8FbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z9GYEgsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAy4zyCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nc3BwMFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IuX5eThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZCWmtu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rvldh2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RJiAIC4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJetYKZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRYXvSnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sg7cBkteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AF0NVmXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjN3FSUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hp6BpSujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQBlr3DnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9IJ1IMieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubuwjPk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIHlYauSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8r7sfSLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2JCOYObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A96xshcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kW0WiVp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAAsL6JRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ga1PMDgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwy2U7dLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEq8DO8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PplluxzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGVzo3WPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGLSiwLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmWrs9AwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tkqa0hpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOsW5KMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjQZwjvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmHIMRrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5g9Xwe4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdHwyJhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ON6KBxiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFXleNoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8V8rm0cCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzFnRu0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VZs5KMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfAhmyyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCNdErewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sL2wOqSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZuwQuXaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qPls7AbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAIowBh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmwqKsv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c18hKuJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klSr2FaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMtc7DB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dk07lU8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmrEifRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFpLrL8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMvcRMhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qkSQbRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLI1Y5CXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1aNcyvKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7Vxd1GOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPdMGlXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJLOjAddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRk9ZoSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxYbwFxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gm26rMdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66IoXeEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yySP1cT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4JOExkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74fvwgDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEHzYuVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubj32RHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13tsIDuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnZpQYxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtiCaAhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgSAd8QuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lieqo5v0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wMqCrNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiOn3bjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvwVhq8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQCpCOsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LopXqkbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVmNt7JhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yh4HiXoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Epb8J57DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNE5L04cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmSaS17aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIbXbswkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VO8ML0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eq2xUaNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMafbwLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkDzQ6QyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d44FzzREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvik2rROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ff6vOrmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1bIdektWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nI1q0xURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55Z0VU7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VU8sNg2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnUnqBM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFJ5xy9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7HEVdDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5SMcDtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MufHvcQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpOZXHqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJOOfV0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unqiI1DHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqam6oF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glWl92NGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVxG97biWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ktl8gYZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Q8CcPK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0uAFNvRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTLrveGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBoyDaH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MPmFmhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzXt2irGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hs0rl3SJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRm3HYm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W45OD2qrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QytoBIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81q98jt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sU2Jct8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckCu7kQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlVIhyQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRqY4W70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRVoAFGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8pstfFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7CZrqPj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZAtssIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdMba5hTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfOLpvozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM7o4Pj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHMFR6SMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsmWvE4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ELc1LMoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96hY1jRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahOsh3xOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uuNEonAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwXbUy1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUnGKKKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hR8kTR6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G48H0JfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ii8038ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxtmCu46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swp1vtN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sfgjTh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUhB2RgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVzLiRczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufech2BfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uv5CfMOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uySHzirtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svoKDYZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkAlcbWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExDSd2VVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8T12j5KWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2P4cmFJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kex8wMemWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func covtWPqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTdD1Rk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6gW98RjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdPkovJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlow8GmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func My5s39e3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezUAgVMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iK5jHjW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lLtvm3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewC7e2DiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C98q0ynxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrJK5hEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ae7vrmKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SG1QoMm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pKkohmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DY8LAs4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jo7Vthq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIxcLZgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKGdogPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4MZcFb2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9IZS4hjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcc5PMiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOx2lUlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYy8MxU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBjxW8CUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hThkDr5CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iz9y3GAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lDcC4LGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwE9tfgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QuBU98k0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uxp6cqDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QE1bvCpnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aLRfQoFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ti4UirFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0udBxO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suqhAG1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WC0Ox3bPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbqclXMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbwEpLE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYdkehYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3qgBNt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvYYzTIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xqxhrTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EO4gPnJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gi4CS7x9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55BhVY25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kp8wEfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgGFW1UhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYfn9lFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPa9Q2WfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UyrhTxZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhzFT0lYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61RfM6tzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnWiI4qwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0E5rDPLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3Xd8fsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqHcSQLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDMjSsQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KwFuAxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYjf1zy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exrI7Ix0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9fEKxr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6wlJPDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wn1L5CTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func siXh9GhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mv6sqHFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtErp4o8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gw4UGfHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyOVBIoCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2k2kUIOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5qIVJydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wH1jqPz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztRMa0euWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnfZcUOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HUcZYtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROUiDGE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Az5Fsm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0uKNba6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxrjUacoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dHrzxv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZQfLerQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTXk2X5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func el5J2ZVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HT8RfUviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AW7s191yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWNjkdlQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4AMozXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNm0rwzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SJG1M7K5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8Utz5aLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8n8bLn8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pxGsVJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8AkUOxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DME03RgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9eaTOkEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qT4DDqBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcdrKMD8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCQ9hnygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLYEWTnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XU6qpQ65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pY9wgEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9eBFpj40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcUWP8IHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptt9sUkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzN905ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQREQDUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a47FQvb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpezRXaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 338PbDTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUquQ5Y3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6S3R2aftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTS5h6qcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBU2w64lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9AIyMLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWPmUdnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQ1gZIXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqaZLVLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNqSs1QuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zUOq4CU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipZL5SqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mF1qgDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhFrQBztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQu8kdqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KLtH3YFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOcWSiSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CcyaEDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFmfnnpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87NewjTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIEOW4GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujDtnRF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiDg31bPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CITzAwXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Esj3KJp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7eN3pHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SajzgKUCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ePkxKoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4iB2rsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeIrbWq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxSR22I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bj00OnuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8mJA5vgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SaUWFsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AlZdQEnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nUDusNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZ1vvdqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rd32WTHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWx7iuTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CybHPg3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIENkET3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzvA6mZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QPyaOmCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KRHSvtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyJfDHoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3sqcYlmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXMJ2lsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEsRFF7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0B5RZBvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cgf2otH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CW7qjE2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbeSm9dmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAQTeSntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRF9A67VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIpwR4HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWERma20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuUW8nLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vU00cXRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ox1f0MoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ek4f931uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASaYEMxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKfMuaT7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEnLmGffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVvWmobaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8TQCMKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kDHUexTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSlisKkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMRGXyhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GWpE9HhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0E2D7aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5R2aT1HoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiSbLQ4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZ8aO2m0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zTDaUmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsBk9NLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n55uB4RhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YbQZJG4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdPneyc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4fbnxNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AW1DbH9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJVbG3xxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOg6bbwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pbS500sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqufDjDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTbMDD9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdglX5lOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tgvq31NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2x005MCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBnrc8N6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VMzajSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqiMwjERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkAIUFEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUnkFqD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mi22k3vQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvK7gBm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glWp5xB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PZ0HRaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vznyfh8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQF1qvcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kc9raktCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95dg9OCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fXhHShe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LDklwJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9u85HRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLaqTlzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BndqJEHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYNMsZzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTRgcc5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VMZILAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Arw2VQgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OYOuoVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZ52Z07mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8GOHXbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeXF4rZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8590Mn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckt93IEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQDnVzf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCQIbeRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXCbUpo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOrR5jakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ti7n9zo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGZ6a3lbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3qlw4RiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zS0A6YlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rq8d6OuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2tHWS48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sssK3CpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5r138GYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4JASYk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUPAl2PPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUrME4osWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uaAPrrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPWr3FMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgBNkKOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5kRlu79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPPWfdYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0730qvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCyWApe4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5szExZy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmoVv87sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDBetMDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiXGEomWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwuqjGsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o94nvgpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htYCvOwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YN6mfwZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjFVbroOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GI0j9KiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hm5zIz2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJMC7sbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sY9073QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39UcBluFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9k1vMuA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTl1QVspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1okt7rbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Szn776lGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQbxO50lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ze2SLewRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtCaWqBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J74MENcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LisFcxi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkGIGNPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsZWPegAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WMr3T4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlFMPtDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lku23joLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZlluINWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGQTmR8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qC6BaXdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPuJXUkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfe0s1AvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMs98zpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byP6wSoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYxycJzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XkTlcD2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xyLTj8yRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NxJD6lsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YE9Ln1JUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpyLMdYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVOHQ5FNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AF7GxeO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cK5K4vH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VCZOSDknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9SASOI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gnFxg6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBcIZUYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cY1WLGQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXa18n06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6FuL8qaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KC976KZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21g2DenpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QhwOII59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVy7zDpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPe3ZsZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDvA8waLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzgBHzvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hiZj8WGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7lV3ur6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJmiwuqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thioVi0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPRWY9RkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39IyZWZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjT2D1fmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUcGcbJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWdhRXwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzm6vKnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Doz74ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIvM0abpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6gmuoNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1aFZMZFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OxT5W3lrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPJ8PhOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v00SD9fdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeGYcGyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAvOR2BUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DjfK7qJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnXUEVOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIjjxS1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k17lYRJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pElud7MeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsTnLAukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tfr8lYLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bjci00uhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdiVYMSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e77UcnenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZ3voKg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q50rxhbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wyjUjUSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTTGswNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9RsmcLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nd6sgzaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9EvPc26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X2WCrZVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUcnTF1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fkhYpXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rfvuoI6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ntgb92e3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYsgF5a7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QNrJhTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jk22pevKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSEo38mzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ml9f2jAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPQMTquvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWdLWQkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LR8I3UmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHzrEasIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7xmMUGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbHM0ZmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xOzIpsLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTLEtHEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvqUSNhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMeZ6JfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84lI35NxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfcneZGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxKu1ZYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFoMFKioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySdhvzfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cw7N6rq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wadcToiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWHxjY3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reonQ7rmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZYUzzKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8xU3UshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J93AOgWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5O3EmelLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cr0tKcitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KuUntTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rpIRaBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYe7yvAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnQiF6fxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PORdRRjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xf6JcjPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIinNtCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQJwUGsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBQJE9aEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hKigyQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWqm6wllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Av4NpsLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKJ3BErJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsClnSdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeX8bstJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0GipUrbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7rBXl70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4lj0FjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbNlIVCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqIWdDVwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goX8O6zFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rsQFDtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5z0P4GfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zW9XQsdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tobs9kUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCEUzhXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wi4whLaNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDP29PSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKKiA6IFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0ebxRVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfYfDaddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCRa7yXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkVDZGEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZVB4Wz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjQ0rncYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNJaJPnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOcX8zCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZET7GhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eb4paD1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wAKq7uuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FG3bqHpOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3XBlmiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGU0j7OwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MXXPED21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4yWDRYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wE2OaDDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mRCM5AJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKjdngwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVqluCZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oT7jSGcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpbUxDozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIgJNB4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hR5fMCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUPbI4hiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoIZnObOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExFV0LezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzUxIFejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HR2oHe3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atyJ6Ot9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTZqwU2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R24NtWw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KiLRPq2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLNP5Uf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acYcI9oDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imF7eE9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDkLMmUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrJeMr51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pu6RGx3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSpq6zSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NutAuRh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CuHaRED0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZwTtFneGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5iDvAotcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9gHNzSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ynu0w3jXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KWzaGQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 839YvE2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAkqPndTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bMh7enC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BtqtarrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWEdEdwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzSV91quWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9fgwvi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7X0Z2ZfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZwe826FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3LaoNAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GG0Kd8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRcSg0LyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvOkf8abWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1mBa4hXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARD2PuOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwobLinSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKt0S6sWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTz4nXVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlZJHd4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzNONSMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkF1BQdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlurAxfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mqfw9pOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEckgDj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgBZph7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQpq9YwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KSYfDCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BAaWGnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5HpFYMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCmuKb94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pEArNP9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLKyeC5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFKwCltGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSCHPM1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZbeTYQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtMVK82IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4n2ozJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mY7pAjYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzBsS1esWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1UXGmmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Tbn5vTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYu63Yk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1LXBVh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDFzg0ErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNbHstj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfGCPXjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmX12dm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6Mocx5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g71uifSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPLGge1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sY6wg0HXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1jDKHOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmuJB6uLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVFBT6vwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKYO5CEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGYbrBQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCZX3TRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxRePuZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSScpdoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cMxcVYMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WifOFDfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icB4D8jxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSg004mUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jutzvvFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlIfOQ0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DReScWVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWLuiQC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wl4zQRHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYqWRkcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ouTL1tUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0H4MfIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftB5dFr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvTAK7iHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CvBbt1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtMtmIORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CavwWdpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dc6VxdmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUBvjOsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUX2ZWsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IH0lapevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p58TiaUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQEkm1WzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4StPLBskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPyMagT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUKgzUj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmUi2Xj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XH8CudQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmPzLND5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ij7FV9WWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxrR8r2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5BFh9dlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QXAs7DaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwz96f3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLk02XLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyIAC4rCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ul621GZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HP3uPeX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FEnehcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVtrbW3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gs3Yw4IbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Hing2ZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3baIGL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJkhFdpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKG3kGyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OV27tBAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRfKSFuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CaIXKqArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nGGe2JKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMgmF70bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9quT1PzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dZqBPLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqbZHrwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtY4KWDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6VfcNsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F31IQL7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxBJYGqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkUvvfdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2jetEk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drX6RQGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3eaSOlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aofhKxkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lzhXsr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMjjCGbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0puCsSLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oR857ETAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bIgaHpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QO1yy7AVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXPRYDetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68ZsFTHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EWpLOh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5T9d3mhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LH72NYodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1FsgroqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4bwYUoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qns8PGXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

