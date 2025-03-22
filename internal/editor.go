package internal

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func IsEditorOpened() (bool, string, error) {
	// windows, linux, darwin (mac os)
	os := runtime.GOOS

	var output []byte
	var err error

	switch os {
	case "windows":
		output, err = exec.Command("tasklist", "/FO", "CSV").Output()
	case "linux", "darwin":
		output, err = exec.Command("ps", "aux").Output()
	default:
		return false, "", fmt.Errorf("unsupported os: %s", os)
	}

	if err != nil {
		return false, "", fmt.Errorf("error executing os command: %v", err)
	}

	outputstr := string(output)

	// Map of editor executable names and their friendly names
	editorMap := map[string]string{
		// VSCode
		"Code.exe": "Visual Studio Code",
		"code":     "Visual Studio Code",
		// JetBrains
		"idea64.exe":     "IntelliJ IDEA",
		"idea":           "IntelliJ IDEA",
		"pycharm64.exe":  "PyCharm",
		"pycharm":        "PyCharm",
		"clion64.exe":    "CLion",
		"clion":          "CLion",
		"goland64.exe":   "GoLand",
		"goland":         "GoLand",
		"webstorm64.exe": "WebStorm",
		"webstorm":       "WebStorm",
		// Terminal editors
		"vim":  "Vim",
		"nvim": "Neovim",
		// GUI editors
		"sublime_text.exe": "Sublime Text",
		"subl":             "Sublime Text",
		"atom.exe":         "Atom",
		"atom":             "Atom",
		"notepad++.exe":    "Notepad++",
		// Emacs
		"emacs.exe": "Emacs",
		"emacs":     "Emacs",
		// Eclipse
		"eclipse.exe": "Eclipse",
		"eclipse":     "Eclipse",
	}

	// Check if any editor is running
	for procName, editorName := range editorMap {
		// On Windows, more precise matching
		if os == "windows" {
			// Look for process name surrounded by quotes or at end of line
			if strings.Contains(outputstr, "\""+procName+"\"") {
				return true, editorName, nil
			}
		} else {
			// On Unix, need to be more careful about partial matches
			lines := strings.Split(outputstr, "\n")
			for _, line := range lines {
				// Check for word boundaries to avoid false positives
				if strings.Contains(line, " "+procName+" ") ||
					strings.HasSuffix(line, " "+procName) ||
					strings.Contains(line, "/"+procName) {
					return true, editorName, nil
				}
			}
		}
	}

	return false, "", nil
}
