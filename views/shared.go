package views

import (
	"context"
	"imaginai/models"
)

func AuthenticatedUser(ctx context.Context) models.AuthenticatedUser {
	if user, ok := ctx.Value(models.UserContextKey).(models.AuthenticatedUser); ok {
		return user
	}

	return models.AuthenticatedUser{}
}
