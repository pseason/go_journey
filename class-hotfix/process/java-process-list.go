package process

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// 26744 DbStart
var r = regexp.MustCompile(`(\d+) ([a-zA-Z0-9.\\:]+)`)

type JavaProcess struct {
	PID  int
	Name string
}

// JavaProcessList 返回当前服务器的Java进程列表
func JavaProcessList() []JavaProcess {
	list := make([]JavaProcess, 0)

	cmd := exec.Command("jps", "-l")
	log.Printf("cmd.Stdout...")
	output := bytes.NewBuffer(nil)
	cmd.Stdout = output
	if err := cmd.Start(); err != nil {
		log.Printf("Start jps failed:%v", err)
		return nil
	}
	log.Printf("Running jps...")
	cmd.Wait()

	text := string(output.Bytes())
	log.Printf("[%s]", text)

	lines := strings.Split(text, "\n")
	for i, line := range lines {
		line = strings.Trim(line, "\r")
		items := r.FindStringSubmatch(line)
		if len(items) > 0 {
			log.Printf("line %d:[%s] items:%v", i, line, items[1:])
			name, processID := items[2], items[1]
			pid, _ := strconv.Atoi(processID)

			// 忽略自己进程
			if name == "sun.tools.jps.Jps" {
				continue
			}

			// 忽略阿里云java进程
			if strings.Index(name, "com.aliyun") != -1 {
				continue
			}

			list = append(list, JavaProcess{
				Name: name,
				PID:  pid,
			})
		} else {
			log.Printf("line %d:[%s] items:%v", i, line, items)
		}
	}

	// 输出格式
	//27808 sun.tools.jps.Jps
	//1508 com.aliyun.tianji.cloudmonitor.Application
	//23078 WorldServer.jar

	log.Printf("Done")

	return list
}
