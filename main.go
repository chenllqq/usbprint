package main

import (
    "fmt"
//    "io"
//    "strings"
    
    //"os"
    //"io/ioutil"
    //"os/exec"
    //"strconv"
    "time"
)

var download bool


func task_download(){
    
    fmt.Println(" task download is running ")

    download = false
    
}

func task_led(){
    var count uint
    var key bool
    var err string

    for true{
        count++
        time.Sleep( time.Millisecond )
        
        //运行灯
        if count == 1000 {
            gpio_set(17, 1)//LED 1
        }else if count == 2000 {
            gpio_set(17, 0)//LED 1
            count = 0
        }
        if download == false{
            //按键检测
            key, err = gpio_read(18)
            /*if err == "" {
                if key == true {
                    download = true
                    go task_download()
                }
            }*/
        }
        
    }
}

func task_usbprint(){
    
    for true{
        time.Sleep(1000*time.Millisecond)
    }
}


func main() {
    
    
    dev_gpio_init()
    dev_usbprint_init()
    
    go task_led()
    go task_usbprint()
    
    fmt.Println("program is running...")
    
    for true{
    }
    return

/*
1. 关键字 defer 用于注册延迟调用。
2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
3. 多个defer语句，按先进后出的方式执行。
4. defer语句中的变量，在defer声明时就决定了。
*/
    //defer iorwc.Close()
    
    //page_pinrt_command( iorwc )
    //pinrt_QRCode( iorwc )
    
   // buffer := make([]byte, 8000 )
    
    //num, err := iorwc.Read(buffer)
    
    /*if err != nil {
        fmt.Println(err)
        fmt.Println("num = ", num)
        return
    }*/
    
    
    /*
    num, err = iorwc.Write( []byte("AT\r\n") )
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
    */
    //打印输出读到的信息
    //fmt.Println(tmpstr)

    return
}
