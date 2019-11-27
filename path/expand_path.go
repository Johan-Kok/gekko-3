package path

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Expand(path string) (string, error) {

	if len(path) == 0 {
		return "", nil
	}

	if strings.HasPrefix(path, "$HOME") || strings.HasPrefix(path, "~") {
		return HomeExpand(path)
	}

	if strings.HasPrefix(path, "$") {
		end := strings.Index(path, string(os.PathSeparator))
		envPath := os.Getenv(path[1:end])
		if len(envPath) == 0 {
			return "", fmt.Errorf("env not set %s", path[1:end])
		}
		return filepath.Join(envPath, path[end:]), nil
	}

	if filepath.IsAbs(path) {
		return filepath.Clean(path), nil
	}

	p, err := filepath.Abs(path)
	if err == nil {
		return filepath.Clean(p), nil
	}
	return path, err
}
