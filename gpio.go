package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "os/exec"
    "strconv"
    "strings"
)


func FileExist(path string) bool {
  _, err := os.Lstat(path)
  return !os.IsNotExist(err)
}

func ExecCommand(strCommand string)(string){

    cmd := exec.Command("/bin/bash", "-c", strCommand)

    stdout, _ := cmd.StdoutPipe()
    if err := cmd.Start(); err != nil{
        fmt.Println("Execute failed when Start:" + err.Error())
        return ""
    }
 
    out_bytes, _ := ioutil.ReadAll(stdout)
    stdout.Close()
 
    if err := cmd.Wait(); err != nil {
        fmt.Println("Execute failed when Wait:" + err.Error())
        return ""
    }
    return string(out_bytes)
}

/*
*设置GPIO的模式：GPIO  输入/输出，值
*num:gpio号
*mode：模式 "input": 输入  "output":输出
*value：仅输出模式时有效，输出值, true: 输出1， false：输出0
*/
func gpio_init(num int, mode string, value int) string {
    var path string
    path = "/sys/class/gpio/gpio"
    //使用strconv.Itoa将int转为string
    //使用strconv.Atoi将string转为int
    path += strconv.Itoa(num) 
    
    var cmd string
    
    if FileExist(path) == false {
        cmd += "echo "
        cmd += strconv.Itoa(num)
        cmd += " > /sys/class/gpio/export"

        ExecCommand( cmd )
    }
    
    
    if mode == "input" {
        cmd = "echo in > "
        cmd += path
        cmd += "/direction"
        
        ExecCommand( cmd )
        //fmt.Println( cmd )
        return ""
    }else{
        cmd = "echo out > "
        cmd += path
        cmd += "/direction"
    }
    //fmt.Println( cmd )
    ExecCommand( cmd )
    
    if value == 0 {
        cmd = "echo 0 > "
    }else{
        cmd = "echo 1 > "
    }
    
    cmd += path
    cmd += "/value"
    ExecCommand( cmd )
    
    return ""
}
/*
 * 读取GPIO值
 * num: GPIO端口号
 * return: vlaue, err_string
 */
func gpio_read(num int) (bool, string) {
    
    var path string
    path = "/sys/class/gpio/gpio"
    path += strconv.Itoa(num)
    
    if FileExist(path) == false {
        return false,string(path + " not exist")
    }
    var cmd string
    
    cmd = "cat "
    cmd += path
    cmd += "/value"
    
    outbyte := ExecCommand( cmd )
    
    //fmt.Println("outbyte==", outbyte)
    //这里低电平有效
    if strings.Contains(outbyte, "1"){
        return false,""
    }
    
    return true, ""
}

/*
 * 设置GPIO值
 * num: GPIO端口号
 * value: 输出值 0 1
 */
func gpio_set(num int, value int) string{
    var path string
    path = "/sys/class/gpio/gpio"
    path += strconv.Itoa(num)
    
    if FileExist(path) == false {
        return string(path + " not exist")
    }
    
    var cmd string
        
    if value == 0 {
        cmd = "echo 0 > "
    }else{
        cmd = "echo 1 > "
    }
    
    cmd += path
    cmd += "/value"
    
    ExecCommand( cmd )
    
    return ""
}

func jtag_sel(num int){
    if num == 0 {
        gpio_set(21, 0)
        gpio_set(20, 1)
        gpio_set(26, 1)
    } else if num == 1 {
        gpio_set(21, 1)
        gpio_set(20, 0)
        gpio_set(26, 1)
    } else if num == 2 {
        gpio_set(21, 1)
        gpio_set(20, 1)
        gpio_set(26, 0)
    }
}

func jtag_disable(){
    gpio_set(21, 1)
    gpio_set(20, 1)
    gpio_set(26, 1)
}

func dev_gpio_init(){
    //led
    gpio_init(23, "out", 1)
    gpio_init(25, "out", 1)
    gpio_init(22, "out", 1)
    gpio_init(24, "out", 1)
    gpio_init(27, "out", 1)
    gpio_init(17, "out", 1)
    
    //key
    gpio_init(18, "input", 1)
    gpio_init(4, "input", 1)
    gpio_init(5, "input", 1)
    gpio_init(6, "input", 1)
    
    //jtag mux ctrl
    gpio_init(26, "out", 1)
    gpio_init(20, "out", 1)
    gpio_init(21, "out", 1)
    
    //sio mux ctrl
    gpio_init(16, "out", 1)
    gpio_init(19, "out", 1)
    gpio_init(13, "out", 1)
    gpio_init(12, "out", 1)
    
    gpio_set(25, 1)//LED 6
    gpio_set(24, 1)//LED 5
    gpio_set(22, 1)//LED 4
    gpio_set(23, 1)//LED 3
    gpio_set(27, 1)//LED 2
    gpio_set(17, 0)//LED 1
    
}