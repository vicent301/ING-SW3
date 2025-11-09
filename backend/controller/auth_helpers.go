package controllers

import (
	"strings"

	"backend/domain"
	"github.com/gin-gonic/gin"
)

func getEmailFromCtx(c *gin.Context) (string, bool) {
	if v, ok := c.Get("email"); ok {
		if s, ok2 := v.(string); ok2 && strings.TrimSpace(s) != "" {
			return strings.TrimSpace(s), true
		}
	}
	return "", false
}

func getUserEmailString(v any) (string, bool) {
	if s, ok := v.(string); ok && strings.TrimSpace(s) != "" {
		return strings.TrimSpace(s), true
	}
	return "", false
}

func getUserEmailUser(u domain.User) (string, bool) {
	if strings.TrimSpace(u.Email) == "" {
		return "", false
	}
	return strings.TrimSpace(u.Email), true
}

func getUserEmailUserPtr(u *domain.User) (string, bool) {
	if u == nil || strings.TrimSpace(u.Email) == "" {
		return "", false
	}
	return strings.TrimSpace(u.Email), true
}

func getUserEmailMap(m map[string]any) (string, bool) {
	if s, ok := m["email"].(string); ok && strings.TrimSpace(s) != "" {
		return strings.TrimSpace(s), true
	}
	return "", false
}
