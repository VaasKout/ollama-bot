package logger

import "log/slog"

func Error(err error) slog.Attr {
	return slog.Any("error", err)
}

func String(key, msg string) slog.Attr {
	return slog.String(key, msg)
}

func Any(key string, data any) slog.Attr {
	return slog.Any(key, data)
}

func Int(key string, data int) slog.Attr {
	return slog.Int(key, data)
}
