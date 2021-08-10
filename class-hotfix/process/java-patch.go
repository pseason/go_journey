package process

import (
	"bytes"
	"fmt"
	"os/exec"
)

// JavaPatch 给指定进程打补丁
// pid 要打补丁的进程ID
// classname 要打补丁的类的名称
// filePath 要打的补丁的文件路径
// java -jar attach.jar 14812 /root/hotswap/agent.jar com.game.hhmf.HhmfTable,/root/hotswap/HhmfTable.class
func JavaPatch(pid int, className, filePath string) (string, error) {
	args := make([]string, 0)
	args = append(args, "-jar")
	args = append(args, "/root/hotswap/attach.jar")
	args = append(args, fmt.Sprintf("%d", pid))
	args = append(args, "/root/hotswap/agent.jar")
	args = append(args, className+","+filePath)
	cmd := exec.Command("java", args...)
	output := bytes.NewBuffer(nil)
	cmd.Stdout = output
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("start patch process  failed:%v", err)
	}

	if err := cmd.Wait(); err != nil {
		return "", fmt.Errorf("run patch process  failed:%v", err)
	}
	return string(output.Bytes()), nil
}
