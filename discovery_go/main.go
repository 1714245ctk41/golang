package main

// import (

// )

import (
    "flag"
    "fmt"
    "os"
    "runtime"
    "time"
    "github.com/chromedp/chromedp"
    "context"
    "log"
)




func main() {
   collycrawl()
}










var n = flag.Int("n", 1e5, "Number of goroutines to create")

var ch = make(chan byte)
var counter = 0

func f() {
    counter++
    <-ch // Block this goroutine
}
func seeComputer(){
    flag.Parse()
    if *n <= 0 {
            fmt.Fprintf(os.Stderr, "invalid number of goroutines")
            os.Exit(1)
    }

    // Limit the number of spare OS threads to just 1
    runtime.GOMAXPROCS(1)

    // Make a copy of MemStats
    var m0 runtime.MemStats
    runtime.ReadMemStats(&m0)

    t0 := time.Now().UnixNano()
    for i := 0; i < *n; i++ {
            go f()
    }
    runtime.Gosched()
    t1 := time.Now().UnixNano()
    runtime.GC()

    // Make a copy of MemStats
    var m1 runtime.MemStats
    runtime.ReadMemStats(&m1)

    if counter != *n {
            fmt.Fprintf(os.Stderr, "failed to begin execution of all goroutines")
            os.Exit(1)
    }

    fmt.Printf("Number of goroutines: %d\n", *n)
    fmt.Printf("Per goroutine:\n")
    fmt.Printf("  Memory: %.2f bytes\n", float64(m1.Sys-m0.Sys)/float64(*n))
    fmt.Printf("  Time:   %f µs\n", float64(t1-t0)/float64(*n)/1e3)
}
  // size := ""
    // html, err := GetHttpHtmlContent("https://www.acfc.com.vn/old-navy-giay-nu-oln-100404-01.html",
    //     `#product-options-wrapper > div > div.swatch-opt`,
    //     `#product-options-wrapper > div > div.swatch-opt`,
    // )
    
    // if err != nil{
    //     fmt.Println("Error: ", err)
    // }
    // fmt.Println(html)

//Get the data crawled from the website
func GetHttpHtmlContent(url string, selector string, sel interface{}) (string, error) {
    options := []chromedp.ExecAllocatorOption{
        chromedp.Flag ("headless", true), // debug
        chromedp.Flag("blink-settings", "imagesEnabled=false"),
        chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
    }
    //Initialization parameters, first pass an empty data    
    options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

    c, _ := chromedp.NewExecAllocator(context.Background(), options...)

    // create context
    chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
    //Execute an empty task and create a chrome instance in advance
    chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

    //Create a context with a timeout of 40s
    timeoutCtx, cancel := context.WithTimeout(chromeCtx, 10*time.Second)
    defer cancel()

    var htmlContent string
    err := chromedp.Run(timeoutCtx,
        chromedp.Navigate(url),
        chromedp.WaitVisible(selector),
        chromedp.OuterHTML(sel, &htmlContent, chromedp.ByJSPath),
    )
    if err != nil {
        fmt.Println("Run err : %v\n", err)
        return "", err
    }
    //log.Println(htmlContent)

    return htmlContent, nil
}


// ctx, cancel := chromedp.NewContext(context.Background())
//     defer cancel()

//     ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
//     defer cancel()

//     var result []string
//     err := chromedp.Run(ctx,
//         chromedp.Navigate(`https://www.acfc.com.vn/old-navy-giay-nu-oln-100404-00.html`),
//         chromedp.WaitVisible(`#product-options-wrapper > div > div.swatch-opt`, chromedp.ByQuery),
//         chromedp.Text(`div.swatch-option`, &result),
//     )
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println(result)