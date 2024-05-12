package crunchyTools

import (
	"fmt"
	"os"
	"strconv"
)

func ForkApp(args []string, processName string, pidToFile bool) {
	cmd := (processName)
	procAttr := new(os.ProcAttr)
	procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	if process, err := os.StartProcess(cmd, args, procAttr); err != nil {
		fmt.Printf("ERROR Unable to run %s: %s\n", cmd, err.Error())
	} else {
		fmt.Printf("DAEMON - %s running as pid %d\n", cmd, process.Pid)

		if pidToFile {
			//write pid to file
			pid := strconv.Itoa(process.Pid)
			errWriting := StringToFile(pid, "pid.txt", 0644)
			HasError(errWriting, "forkItSelf - StringToFile", false)
		}
	}
}
