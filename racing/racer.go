package racing

import (
	"net/http"
    "fmt"
	// "time"
)


func WebRacer(url1, url2 string) (winner string) {

    c1 := make(chan string)
    c2 := make(chan string)

    go ping2(url1, c1)
    go ping2(url2, c2)

    select {
    case <-c1:
        return url1
    case <-c2:
        return url2
    }


    // select {
    // case <-ping(url1):
    //     return url1
    // case <-ping(url2):
    //     return url2   
    // }
}

func ping2(url string, c chan string) {
    r, err := http.Get(url)
    fmt.Println(r.Status)
    if err != nil {
        c <- r.Status
    }
    close(c)
}

// func ping(url string) chan string {

//     c := make(chan string)

//     go func() {
//         http.Get(url)
//         c <- url
//     }()

//     close(c)
//     return c
// }

// func WebRacer(url1, url2 string) (winner string) {
//     duration1 := getDuration(url1)
//     duration2 := getDuration(url2)

//     if duration1 > duration2 {
//         return url2
//     }

//     return url1
// }

// func getDuration(url string) time.Duration {
//     start := time.Now()
//     http.Get(url)
//     return time.Since(start)
// }