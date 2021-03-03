package main

import (
    "fmt"
    "io"
    "time"
    "github.com/tarm/goserial"
)


func read_mac( sio io.ReadWriteCloser ){
    
    sbuf := make([]byte, 8 )
    
    var slen uint = 0
    
    //初始化打印机
    sbuf[slen] = 0x01
    slen++
    sbuf[slen] = 0x03
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x00
    slen++
    sbuf[slen] = 0x04
    slen++
    sbuf[slen] = 0x44
    slen++
    sbuf[slen] = 0x09
    
    num, err := sio.Write( sbuf )
    if err != nil {
        fmt.Println(err)
        return
    }else{
        fmt.Println("send ok, num = ", num)
    }
}

func dev_rs485_call_mac_addr() (string,string) {

    cfg := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600 , ReadTimeout:3 }
    
    iorwc, err := serial.OpenPort(cfg)
    
    if( err != nil ){
        fmt.Println(err)
        return "","can not open tty!"
    }

    buffer := make([]byte, 100 )
    
    read_mac(iorwc)
    
    time.Sleep(200*time.Millisecond)
    
    num, err := iorwc.Read(buffer)
    
    if err != nil {
        fmt.Println(err)
    }
    var tmpstr string
    
    if num > 0 {
        tmpstr += fmt.Sprintf("%x", string(buffer[:num]))
    }else{
        return  "FFFFFFFFFFFFFFFFFF","read num is 0"
    }
    
    fmt.Println(tmpstr)
    
    iorwc.Close()
    
    var macaddr string
    macaddr += fmt.Sprintf("%X", string(buffer[3:11]))
    fmt.Println(macaddr)

    return  macaddr,""
}