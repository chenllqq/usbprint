package main

import (
    "fmt"
    "strings"
    "github.com/tarm/goserial"
)

func main() {
    cfg := &serial.Config{Name: "/dev/ttyUSB0", Baud: 115200 , ReadTimeout:3 }
    
    iorwc, err := serial.OpenPort(cfg)
    
    if( err != nil ){
        fmt.Println(err)
        return
    }
    
    defer iorwc.Close()

    buffer := make([]byte, 8000 )

    num, err := iorwc.Read(buffer)

    /*if err != nil {
        fmt.Println(err)
        fmt.Println("num = ", num)
        return
    }*/
    
    num, err = iorwc.Write([]byte("AT\r\n"))
    if err != nil {
        fmt.Println(err)
        return
    }
    
    var tmpstr string = ""

    for i := 0; i < 3000; i++ {
        num, err = iorwc.Read(buffer)
        if num > 0 {
            tmpstr += fmt.Sprintf("%s", string(buffer[:num]))
        }

        //查找读到信息的结尾标志
        if strings.LastIndex(tmpstr, "\r\nOK\r\n") > 0 {
            break
        }
    }

    //打印输出读到的信息
    fmt.Println(tmpstr)

    return
}
