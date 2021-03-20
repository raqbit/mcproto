package mcproto

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
}

type noopLogger struct {
}

func (n noopLogger) Debug(_ ...interface{}) {
}

func (n noopLogger) Debugf(_ string, _ ...interface{}) {
}
