package platform

type Logger interface {
	Info(string)
	Error(string)
}
