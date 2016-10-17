package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/go-template/html"
    "strconv"
    "fmt"
    "time"
)

type Nyan struct {
    Num int
    Con iris.WebsocketConnection
}

var cats []*Nyan

// SETTINGS
const TIMEOUT = 1250
const IMGSRC = "nyan.png"

func main() {
    cats = make([]*Nyan, 0)

    iris.UseTemplate(html.New())

    iris.Config.Websocket.Endpoint = "/nyansocket"

    nyanNum := 0
    bckColor := "FFFFFF"
    if IMGSRC == "nyan.png" {
        bckColor = "003466"
    }

    iris.Get("/", func (ctx *iris.Context) {
        nyanNum++
        ctx.RenderWithStatus(iris.StatusOK, "nyan.html",
            map[string]interface{}{
                "Number": nyanNum,
                "Host": ctx.HostString(),
                "Timeout": TIMEOUT,
                "Duration": TIMEOUT*2,
                "ImgSrc": IMGSRC,
                "BackgroundColor": bckColor,
            })
    })

    iris.Get("/" + IMGSRC, func (ctx *iris.Context) {
        ctx.ServeFile(IMGSRC, false)
    })

    iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
        c.Join("nyan")

        c.On("nyan", func (msg string) {
            n, _ := strconv.Atoi(msg)
            cats = append(cats, &Nyan{Con: c, Num: n});
            fmt.Println("Client connected as #" + msg)
        })
    })

    go iris.Listen(":8080")
    defer iris.Close()

    fmt.Println("Server started! Press enter to start the madness...")
    fmt.Scanln()
    fmt.Println("NYANYANYANYANYANYAN!!1! (Clients: " + strconv.Itoa(len(cats)) + ")")

    for true {
        for i, cat := range cats {
            fmt.Println("Nyaning at: " + strconv.Itoa(i))
            cat.Con.Emit("nyan", "go");
            time.Sleep(time.Duration(TIMEOUT) * time.Millisecond)
        }
    }
}