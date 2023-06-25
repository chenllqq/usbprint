package main

import (
    "fmt"
    "strings"
    "time"
)

var Version string

var download bool

/*
 *功能：从串口读取设备的唯一ID，并通过打印机打印出二维码
 */
func task_printf_code(){
    
    macaddr,err:= dev_rs485_call_mac_addr()
    if err != "" {
        fmt.Println(err)
        //点串口读取失败告警灯
        gpio_set(24, 0)//LED 1
        
        return
    }
    
    fmt.Println(macaddr)
    
    macaddr += " "
    macaddr += time.Now().Format("2006-01-02 15:04:05")
    dev_usbprint_printf(macaddr)
}

/*
 *功能：烧写任务
 */
func task_download(){
    
    fmt.Println("task download stm32 is running ")
    
    //清除示灯
    gpio_set(23, 1)//LED 3
    gpio_set(22, 1)//LED 4
    gpio_set(24, 1)//LED 5
    
    //点下载灯
    gpio_set(27, 0)//LED 2
    time.Sleep(1000*time.Millisecond)

    //JTAG 复用控制
    jtag_sel(0)

    //调用下载脚本
    var cmd string
    cmd = "sh /home/pi/Work/download.sh"
    outbyte := ExecCommand(cmd)
    
    fmt.Println(outbyte)
    
    if  strings.Contains( outbyte, "Program & Verify speed:")||
        strings.Contains( outbyte, "Contents already match"){
        fmt.Println("============download successful")
        gpio_set(23, 0)//LED 1
        
        //调用打印功能
        go task_printf_code()
    }else{
        fmt.Println("============download fail")
        gpio_set(22, 0)//LED 1
    }
    
    fmt.Println("task download stm32 complete ")
    //关闭下载灯
    gpio_set(27, 1)//LED 1
    
    //关闭jtag
    jtag_disable()
    
    download = false
}

/*
 * 烧写efr32任务
 */
func task_download_efr32(){
    fmt.Println("task download efr32 is running ")
    
    //清楚下载指示灯
    gpio_set(23, 1)//LED 1
    gpio_set(22, 1)//LED 1
    
    //点下载灯
    gpio_set(27, 0)//LED 1
    time.Sleep(1000*time.Millisecond)

    //JTAG 复用控制
    jtag_sel(2)

    //调用下载脚本
    var cmd string
    cmd = "sh /home/pi/Work/download1.sh"
    outbyte := ExecCommand(cmd)
    
    fmt.Println(outbyte)
    
    if  strings.Contains( outbyte, "Program & Verify speed:")||
        strings.Contains( outbyte, "Contents already match"){
        fmt.Println("============download successful")
        gpio_set(23, 0)//LED 1
    }else{
        fmt.Println("============download fail")
        gpio_set(22, 0)//LED 1
    }
    
    fmt.Println("task download efr32 complete ")
    //关闭下载灯
    gpio_set(27, 1)//LED 1
    
    //关闭jtag
    jtag_disable()
    
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
        if count == 500 {
            gpio_set(17, 1)//LED 1
        }else if count == 1000 {
            gpio_set(17, 0)//LED 1
            count = 0
        }
        
        if (download == false)&&(count%50==0){
            //按键检测
            key, err = gpio_read(18)
            if err == "" {
                if key == true {
                    download = true
                    go task_download()
                }
            }
            
            key,err = gpio_read(4)
            if err == "" {
                if key == true {
                    download = true
                    go task_download_efr32()
                }
            }
        }
    }
}


func main() {
    
    Version = "1.0.0"
    
    dev_gpio_init()
    
    go task_led()
    
    fmt.Println("program is running...")
    
    for true{
        time.Sleep( 1000*time.Millisecond )
    }

    
    return
}
