package clipboard

import (
	"fmt"
	"os/exec"
)

// write to clipboard
func Write(content []byte) error {

	cmd := exec.Command("xclip", "-i", "-selection", "clipboard")

	//创建获取命令输出管道
	in, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return err
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return err
	}

	//读取所有输出
	_, err = in.Write(content)
	if err != nil {
		fmt.Println("WriteAll Stdin:", err.Error())
		return err
	}

	if err := in.Close(); err != nil {
		fmt.Println("close:", err.Error())
		return err
	}

	if err := in.Close(); err != nil {
		fmt.Println("close:", err.Error())
		return err
	}
	return nil
}
