package finder

import "fmt"

type TodoItem struct {
	LineNum  int64
	Text     string
	FilePath string
}

func (todo *TodoItem) ToString() string {
	return fmt.Sprintf("%s L:%v:%s", todo.FilePath, todo.LineNum, todo.Text)
}

type Finder interface {
	Find(initialPath string) (todos []*TodoItem, err error)
	FindFromDir(dirPath string) (todos []*TodoItem, err error)
	FindFromFile(filePath string) (todos []*TodoItem, err error)
	FindFromLine(line []byte) (todo string, err error)
}
