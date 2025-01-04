package middleware

import (
	"errors"
	"github.com/charmbracelet/ssh"
	"go.uber.org/zap"
	"keeper/internal/app"
	"keeper/internal/service/user"
	"keeper/pkg/hash"
)

func SignIn(ka *app.App, logger *zap.SugaredLogger) func(ctx ssh.Context, key ssh.PublicKey) bool {
	return func(ctx ssh.Context, key ssh.PublicKey) bool {
		userID, err := ka.UserService.GetUserIDByKey(ctx, hash.Hash(key.Marshal()))
		if err != nil && errors.Is(err, user.ErrUnauthorizedUser) {
			return true
		}

		if err != nil {
			logger.Errorf("unable to auth user due to internal error: %v", err)
			return false
		}

		setUserIDToContext(ctx, userID)
		return true
	}
}

func setUserIDToContext(ctx ssh.Context, userID uint64) {
	ctx.SetValue(user.UserIDContextKey, userID)
}
