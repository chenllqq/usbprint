package main

import (
    "fmt"
    "io"
//    "strings"
    
    //"os"
    //"io/ioutil"
    //"os/exec"
    //"strconv"
    //"time"
    
    "github.com/tarm/goserial"
)

func page_pinrt_command( sio io.ReadWriteCloser ){
    
    sbuf := make([]byte, 2000 )
    
    var slen uint = 0
    
    //初始化打印机
    slen++
    sbuf[slen] = 0x1B
    slen++
    sbuf[slen] = 0x40
    
    //页开始
    slen++
    sbuf[slen] = 0x1a
    slen++
    sbuf[slen] = 0x5b
    slen++
    sbuf[slen] = 0x01
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x80
    slen++
    sbuf[slen] = 0x01
    slen++
    sbuf[slen] = 0xF0
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    
    //页面数据
    slen++
    sbuf[slen] = 0x1A
    slen++
    sbuf[slen] = 0x54
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x10
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x10
    slen++
    sbuf[slen] = 0x00
    
    slen++
    sbuf[slen] = 0xC9
    slen++
    sbuf[slen] = 0xEE
    slen++
    sbuf[slen] = 0xDB
    slen++
    sbuf[slen] = 0xDA
    slen++
    sbuf[slen] = 0xB8
    slen++
    sbuf[slen] = 0xE8
    slen++
    sbuf[slen] = 0xD2
    slen++
    sbuf[slen] = 0xED
    slen++
    sbuf[slen] = 0xBF
    slen++
    sbuf[slen] = 0xC6
    slen++
    sbuf[slen] = 0xBC
    slen++
    sbuf[slen] = 0xBC
    slen++
    sbuf[slen] = 0xD3
    slen++
    sbuf[slen] = 0xD0
    slen++
    sbuf[slen] = 0xCF
    slen++
    sbuf[slen] = 0xDE
    slen++
    sbuf[slen] = 0xB9
    slen++
    sbuf[slen] = 0xAB
    slen++
    sbuf[slen] = 0xCB
    slen++
    sbuf[slen] = 0xBE

    //页面结束
    slen++
    sbuf[slen] = 0x1a
    slen++
    sbuf[slen] = 0x5d
    slen++
    sbuf[slen] = 0x00
    
    //页打印
    slen++
    sbuf[slen] = 0x1a
    slen++
    sbuf[slen] = 0x4f
    slen++
    sbuf[slen] = 0x00
    
    num, err := sio.Write( sbuf )
    if err != nil {
        fmt.Println(err)
        return
    }else{
        fmt.Println("send ok, num = ", num)
    }
}


/*
功能：打印二维码
*/
func pinrt_QRCode( sio io.ReadWriteCloser ,str string){
    
    sbuf := make([]byte, 2000 )
    
    var slen uint = 0
    
    //初始化打印机
    slen++
    sbuf[slen] = 0x1B
    slen++
    sbuf[slen] = 0x40
    
    //页开始
    slen++
    sbuf[slen] = 0x1a
    slen++
    sbuf[slen] = 0x5b
    slen++
    sbuf[slen] = 0x01
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0xc8//页面宽度 25mm 200个点
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x78//页面高度 15mm 120个点
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    
    //页面数据
    slen++
    sbuf[slen] = 0x1A
    slen++
    sbuf[slen] = 0x31
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x02 //版本 [0~20]
    slen++
    sbuf[slen] = 0x01 //ecc [1~4]=[L M Q H]
    slen++
    sbuf[slen] = 0x00 //左上角坐标x 低8位
    slen++
    sbuf[slen] = 0x00 //左上角坐标x 高8位
    slen++
    sbuf[slen] = 0x05 //左上角坐标y
    slen++
    sbuf[slen] = 0x00 //左上角坐标y
    slen++
    sbuf[slen] = 0x04 //二维码大小
    slen++
    sbuf[slen] = 0x00
    
    //二维码数据,最大47个字母
    //str := string("ABCDEFHIJKABCDEFHIJKABCDEFHIJKABCDEFHIJKABCDEFH")
    
    var data []byte
    
    data = []byte(str)
    
    var i int = 0
    for i = 0; i < len(data); i++{
        slen++
        sbuf[slen] = data[i]
    }
    slen++
    sbuf[slen] = 0x00
    
    //页面结束
    slen++
    sbuf[slen] = 0x1a
    slen++
    sbuf[slen] = 0x5d
    slen++
    sbuf[slen] = 0x00
    
    //页打印
    slen++
    sbuf[slen] = 0x1a
    slen++
    sbuf[slen] = 0x4f
    slen++
    sbuf[slen] = 0x00
    
    num, err := sio.Write( sbuf )
    if err != nil {
        fmt.Println(err)
        return
    }else{
        fmt.Println("send ok, num = ", num)
    }
}

/*
功能：走纸功能
*/
func pinrt_null(){

}


func dev_usbprint_printf(str string){

    cfg := &serial.Config{Name: "/dev/ttyUSB1", Baud: 115200 , ReadTimeout:3 }
    
    iorwc, err := serial.OpenPort(cfg)
    
    if( err != nil ){
        fmt.Println(err)
        return
    }
    
    pinrt_QRCode(iorwc, str)
    
    iorwc.Close()
}
