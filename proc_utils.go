package linuxlens

import (
	"io/ioutil"
	"log"
	"strconv"
)

func isNumeric(dirName string) bool {
	_, err := strconv.Atoi(dirName)
	return err == nil
}

func GetProcesses() []LensProcess {
	procPath := "/proc"

	files, err := ioutil.ReadDir(procPath)
	if err != nil {
		log.Fatal(err)
	}

	processes := make([]LensProcess, 0, 1024)
	for _, file := range files {
		if file.IsDir() && isNumeric(file.Name()) {

			pid := file.Name()

			cmdLine, err := parseCmdLine(pid)
			if err != nil {
				continue
			}

			user, err := parseUsername(pid)
			if err != nil {
				continue
			}

			status, err := processState(pid)
			if err != nil {
				continue
			}

			process, err := NewProcess(cmdLine, *user, status)
			if err != nil {
				continue
			}

			processes = append(processes, *process)
		}
	}

	return processes
}
