package linuxlens

import (
	"os/user"
	"strconv"
	"fmt"
)

type LensProcess struct {
	CommandLine string
	UserName    string
	UserID      uint16
	Status      string
}

func NewProcess(commandline string, user user.User, status string) (*LensProcess, error) {

	uid, err := strconv.Atoi(user.Uid)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %s", user.Uid)
	}

	return &LensProcess{
		CommandLine: commandline,
		UserName:    user.Name,
		UserID:      uint16(uid),
		Status:      status,
	}, nil
}