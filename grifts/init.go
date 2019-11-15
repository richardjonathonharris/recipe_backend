package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/richardjonathonharris/recipe_backend/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
