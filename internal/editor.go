package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func IsEditorOpened() (bool, error) {

	// windows, linux, darwin (mac os)
	os := runtime.GOOS

	var output []byte
	var err error

	switch os {
	case "windows":
		output, err = exec.Command("tasklist").Output()
	case "linux", "darwin":
		output, err = exec.Command("ps", "aux").Output()
	default:
		return false, fmt.Errorf("unsupported os")
	}

	if err != nil {
		return false, fmt.Errorf("error executing os command")
	}

	editors := []string{
		"Code.exe", "code", // VSCode
		"idea64.exe", "idea", // JetBrains IntelliJ
		"pycharm64.exe", "pycharm", // JetBrains PyCharm
		"clion64.exe", "clion", // JetBrains CLion
		"vim", "nvim", // Vim & Neovim
		"sublime_text.exe", "subl", // Sublime Text
		"atom.exe", "atom", // Atom
		"notepad++.exe",      // Notepad++
		"emacs.exe", "emacs", // Emacs
		"eclipse.exe", "eclipse", // Eclipse
	}

	outputstr := string(output)

	for _, editor := range editors {
		if strings.Contains(outputstr, editor) {
			return true, nil
		}
	}
	return false, nil
}
