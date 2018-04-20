package model

import (
	"fmt"
)

// BasePage is a base model.
type BasePage struct {
	Title string
}

// String is used for model toString.
func (model *BasePage) String() string {
	return fmt.Sprintf(
		"<Title:%s>",
		model.Title)
}
