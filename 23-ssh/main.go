package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	// 配置 SSH 客户端
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("123456-Ab@"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不建议在生产环境中使用
	}

	// 连接到 SSH 服务器
	host := "localhost:22"
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		fmt.Printf("Failed to dial: %s\n", err)
		return
	}
	defer client.Close()

	// 创建一个新的会话
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s\n", err)
		return
	}
	defer session.Close()

	// 设置伪终端
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // 禁用回显
		ssh.TTY_OP_ISPEED: 14400, // 输入速度
		ssh.TTY_OP_OSPEED: 14400, // 输出速度
	}

	// 请求伪终端
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		fmt.Printf("Request for pseudo terminal failed: %s\n", err)
		return
	}

	// 将标准输入和输出连接到会话
	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// 启动 shell
	if err := session.Shell(); err != nil {
		fmt.Printf("Failed to start shell: %s\n", err)
		return
	}

	// 等待会话结束
	session.Wait()
}
