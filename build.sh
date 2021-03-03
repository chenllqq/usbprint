GOARM=7 GOARCH=arm GOOS=linux go build main.go gpio.go usbprint.go rs485.go

scp main pi@192.168.1.145:/home/pi/Work

