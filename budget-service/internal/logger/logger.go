package logger

// Bu dosya dışarıya açılan unified log API'dir

var impl Implementation

// Implementation arayüzünü tanımlarız
type Implementation interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Fatal(msg string, fields ...interface{})
}

// DI fonksiyonu – zap, logrus, vs. inject edilebilir
func InitLogger(i Implementation) {
	impl = i
}

func Info(msg string, fields ...interface{}) {
	impl.Info(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	impl.Error(msg, fields...)
}

func Fatal(msg string, fields ...interface{}) {
	impl.Fatal(msg, fields...)
}
