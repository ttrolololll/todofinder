package finder

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type TodoFinderImpl struct{}

func (impl *TodoFinderImpl) Find(initialPath string) (todos []*TodoItem, err error) {
	// Check if path is valid
	fi, err := os.Stat(initialPath)
	if err != nil {
		return todos, err
	}

	// Use abs path
	absPath, err := filepath.Abs(initialPath)
	if err != nil {
		return todos, err
	}

	// Process
	if fi.IsDir() {
		todoRes, err := impl.FindFromDir(absPath)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todoRes...)
	} else {
		todoRes, err := impl.FindFromFile(absPath)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todoRes...)
	}

	return todos, nil
}

func (impl *TodoFinderImpl) FindFromDir(dirPath string) (todos []*TodoItem, err error) {
	// TODO: exlusion to use flag/config file
	if strings.Contains(dirPath, "node_modules") {
		return todos, nil
	}

	dirItems, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return todos, err
	}

	for _, item := range dirItems {
		itemPath := filepath.Join(dirPath, item.Name())

		if item.IsDir() {
			todoRes, err := impl.FindFromDir(itemPath)
			if err != nil {
				continue
			}
			todos = append(todos, todoRes...)
		} else {
			todoRes, err := impl.FindFromFile(itemPath)
			if err != nil {
				continue
			}
			todos = append(todos, todoRes...)
		}
	}

	return todos, nil
}

func (impl *TodoFinderImpl) FindFromFile(filePath string) (todos []*TodoItem, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return todos, err
	}
	defer file.Close()

	lineNum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineNum++
		lineContent := scanner.Bytes()

		todo, err := impl.FindFromLine(lineContent)
		if err != nil {
			continue
		}
		if todo == "" {
			continue
		}

		todos = append(todos, &TodoItem{
			LineNum:  int64(lineNum),
			Text:     todo,
			FilePath: filePath,
		})

	}

	return todos, nil
}

func (impl *TodoFinderImpl) FindFromLine(line []byte) (todo string, err error) {
	matched, err := regexp.Match(`(?i)\/\/\s*(todo)`, line)

	if err != nil {
		return "", err
	}

	if !matched {
		return "", nil
	}

	var snippet string
	lineText := strings.TrimSpace(string(line))

	if len(lineText) > 50 {
		snippet = lineText[:50] + "..."
	} else {
		snippet = lineText
	}

	return snippet, nil
}
