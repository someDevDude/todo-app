package api

import (
	"testing"

	"github.com/someDevDude/todo-server/models"
)

func TestValidateTodo_success(t *testing.T) {
	todo := &models.TodoFull{1, "yes", "yes", false}

	errorMessage := ValidateTodo(todo)

	if errorMessage != "" {
		t.Error("errorMessage is not null")
	}

}

func TestValidateTodo_fail_noTitle(t *testing.T) {
	expectedErrorMessage := "Missing title for todo"
	todo := &models.TodoFull{1, "", "yes", false}

	errorMessage := ValidateTodo(todo)

	if errorMessage != expectedErrorMessage {
		t.Error("errorMessage is not correct, expecting " + expectedErrorMessage)
	}

}
