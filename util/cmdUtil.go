package util

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Cmd() {
	cmd := exec.Command("cmd")
	// cmd := exec.Command("powershell")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	go func() {
		// start stop restart
		in.WriteString("cmd /c start C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe") //写入你的命令，可以有多行，"\n"表示回车
	}()
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cmd.Args)
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	rt := out.String() //mahonia.NewDecoder("gbk").ConvertString(out.String()) //
	fmt.Println(rt)

	if strings.ContainsAny(rt, "成功") {
		log.Println("操作成功")
	} else {
		log.Println(rt)
	}
}
