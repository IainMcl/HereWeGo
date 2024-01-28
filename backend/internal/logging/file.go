package logging

import (
	"fmt"
	"time"

	"github.com/IainMcl/HereWeGo/internal/settings"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", settings.AppSettings.RuntimeRootPath, settings.AppSettings.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		settings.AppSettings.LogSaveName,
		time.Now().Format(settings.AppSettings.TimeFormat),
		settings.AppSettings.LogFileExt,
	)
}
