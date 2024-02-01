package util

import "github.com/IainMcl/HereWeGo/internal/settings"

func Setup() {
	JwtSecret = []byte(settings.AppSettings.JwtSecret)
}
