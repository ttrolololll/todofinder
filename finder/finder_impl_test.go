package finder

import "testing"

// For purpose of intent, omitting the other tests

func TestTodoFinderImpl_FindFromLine(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name       string
		pLine      []byte
		expectTodo bool
	}{
		{
			name:       "standard case",
			pLine:      []byte("// TODO: must play dota"),
			expectTodo: true,
		},
		{
			name:       "mixed case",
			pLine:      []byte("// tODo: must play dota"),
			expectTodo: true,
		},
		{
			name:       "nil space // + todo",
			pLine:      []byte("//TODO: must play dota"),
			expectTodo: true,
		},
		{
			name:       "todo at eol",
			pLine:      []byte("a := b + c // TODO: must play dota"),
			expectTodo: true,
		},
		{
			name:       "no todo",
			pLine:      []byte("a := b + c // must play dota TODO: this will not be counted as todo"),
			expectTodo: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			todoFinder := &TodoFinderImpl{}
			todo, _ := todoFinder.FindFromLine(tc.pLine)

			if tc.expectTodo && todo == "" {
				t.Error("expected todo, but empty")
			}

			if !tc.expectTodo && todo != "" {
				t.Error("unexpected todo")
			}
		})
	}
}
