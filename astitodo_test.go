package astitodo_test

import (
	"testing"

	"github.com/asticode/go-astitodo"
	"github.com/stretchr/testify/assert"
)

func TestProcessPath(t *testing.T) {
	todos, err := astitodo.Extract("tests", "tests/excluded.go")
	assert.NoError(t, err)
	assert.Len(t, todos, 6)
	assert.Equal(t, &astitodo.TODO{Line: 5, Message: []string{"Here is a", "multi line todo"}, Filename: "tests/level1/level2.go"}, todos[0])
	assert.Equal(t, &astitodo.TODO{Assignee: "my.weird-email_address+1@email.com", Line: 11, Message: []string{"This is a named TODO"}, Filename: "tests/level1/level2.go"}, todos[1])
	assert.Equal(t, &astitodo.TODO{Assignee: "quentin renard", Line: 16, Message: []string{"Here is another", "multi line todo"}, Filename: "tests/level1/level2.go"}, todos[2])
	assert.Equal(t, &astitodo.TODO{Line: 8, Message: []string{"Is it really your second function?"}, Filename: "tests/level1.go"}, todos[3])
	assert.Equal(t, &astitodo.TODO{Line: 11, Message: []string{"This is a tabbed TODO"}, Filename: "tests/level1.go"}, todos[4])
	assert.Equal(t, &astitodo.TODO{Line: 12, Message: []string{"this a second todo in the same comment group"}, Filename: "tests/level1.go"}, todos[5])
}