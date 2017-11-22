package linuxlens

import (
	"io/ioutil"
	"strings"
	"bytes"
	"fmt"
	"os/user"
	"os"
	"bufio"
	"path/filepath"
)

func resolvePath(pid string, filename string) string {
	return filepath.Join("/proc", pid, filename)
}

func sanitize(data []byte) string {
	data = bytes.Replace(data, []byte("\u0000"), []byte(" "), -1)
	return strings.TrimSpace(string(data))
}

func parseCmdLine(pid string) (string, error) {

	cmdlinePath := resolvePath(pid, "cmdline")

	data, err := ioutil.ReadFile(cmdlinePath)
	if err != nil {
		return "", err
	}

	content := sanitize(data)

	if len(content) == 0 {
		return "", fmt.Errorf("empty pid")
	}

	return strings.TrimSpace(content), nil
}

func parseUsername(pid string) (*user.User, error) {

	loginFile := resolvePath(pid, "loginuid")

	data, err := ioutil.ReadFile(loginFile)
	if err != nil {
		return nil, err
	}

	uid := string(data)

	u, err := user.LookupId(uid)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func processState(pid string) (string, error) {

	statusFile := resolvePath(pid, "status")

	file, err := os.Open(statusFile)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "State:") {
			fields := strings.SplitN(line, "\t", 2)
			if len(fields) >= 2 {
				return fields[1], nil
			}
		}
	}

	return "", nil
}
