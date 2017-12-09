package linuxlens

import (
	"io/ioutil"
	log "github.com/Sirupsen/logrus"
	"strconv"
	"strings"
)

func isNumeric(dirName string) bool {
	_, err := strconv.Atoi(dirName)
	return err == nil
}

func GetProcesses(cfg *ServerConfig) []LensProcess {
	procPath := "/proc"

	files, err := ioutil.ReadDir(procPath)
	if err != nil {
		log.Fatal(err)
	}

	processes := make([]LensProcess, 0, 1024)
	blacklist := cfg.BlacklistPatterns

	for _, file := range files {
		if file.IsDir() && isNumeric(file.Name()) {

			pid := file.Name()

			cmdLine, err := parseCmdLine(pid)
			if err != nil {
				continue
			}

			blacklisted := false
			for _, pattern := range blacklist {
				if strings.Contains(cmdLine, pattern) {
					blacklisted = true
					break
				}
			}

			if blacklisted {
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
