package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
	"strings"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"backend/app/model"
	"backend/app/utils"
)

type Config struct {
	Filename        string
	SymlinkName     string
	AddSource       bool
	Level           int
	TimeFormat      string
	RetentionInDays int64
	RotateInHour    int64
}

func LoggerWithTimeStamp(message string, args ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	if len(args) > 0 {
		if !strings.Contains(message, "%") {
			message = message + fmt.Sprint(args...)
		} else {
			message = fmt.Sprintf(message, args...)
		}
	}
	fmt.Println("INFO:", currentTime, message)
}

func Level(i int) slog.Level {
	return []slog.Level{slog.LevelDebug, slog.LevelInfo}[i]
}

func Initialize(cfg model.CONFIG) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	optrl := []rotatelogs.Option{
		rotatelogs.WithMaxAge(time.Duration(cfg.LOG.RETENTION_IN_DAYS) * 24 * time.Hour),
		rotatelogs.WithRotationTime(time.Duration(cfg.LOG.ROTATE_IN_HOUR) * time.Hour),
		rotatelogs.WithClock(rotatelogs.Local),
	}
	if len(cfg.LOG.SYMINK_NAME) != 0 {
		optrl = append(optrl, rotatelogs.WithLinkName(cfg.LOG.SYMINK_NAME))
	}

	rl, err := rotatelogs.New(

		cfg.LOG.FILE_NAME, optrl...)
	if err != nil {
		panic(err)
	}

	multi := io.MultiWriter(os.Stdout, rl)

	opts := &slog.HandlerOptions{
		AddSource: cfg.LOG.ADD_SOURCE,
		Level:     Level(cfg.LOG.LEVEL),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {

			switch a.Key {
			case slog.TimeKey:
				if len(cfg.LOG.TIME_FORMAT) != 0 {
					ts := a.Value.Time()
					return slog.String("timestamp", ts.In(loc).Format(cfg.LOG.TIME_FORMAT))
				}
			case slog.LevelKey:
				if a.Value.Any().(slog.Level) == slog.LevelInfo {
					return slog.Attr{}
				}
			case slog.MessageKey:
				if len(a.Value.String()) == 0 {
					return slog.Attr{}
				}
			}
			return a
		},
	}
	utils.LOAD_ENV()
	if os.Getenv("LOGGER_STATUS") == "DISABLED" {
		// fmt.Println("Logger Disabled")
	} else {
		json := slog.New(slog.NewJSONHandler(multi, opts))
		slog.SetDefault(json)
	}
}