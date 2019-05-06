package docker

import (
	"fmt"
	"os"
	"strings"

	"github.com/talal/go-bits/color"
)

// Host returns the current Docker Host (only remote).
func Host() string {
	if os.Getenv("CRISP_DOCKER") == "0" {
		return ""
	}
	host := os.Getenv("DOCKER_HOST")
	if host == "" {
		return ""
	} else if strings.Contains(host, "localhost") || strings.Contains(host, "docker.sock") {
		return ""
	}

	return fmt.Sprintf(" 🐋 %v", color.Sprintf(color.Cyan, host))
}
