package finder

import "testing"

func TestTodoItem_ToString(t *testing.T) {
	t.Parallel()

	item := &TodoItem{
		LineNum:  99,
		Text:     "// TODO: must do",
		FilePath: "~/myproject/file.go",
	}

	str := item.ToString()

	if str != "~/myproject/file.go L:99:// TODO: must do" {
		t.Error("output string mismatch")
	}
}
