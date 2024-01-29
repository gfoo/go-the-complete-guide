package iomanager

type IOMAnager interface {
	ReadLines() ([]string, error)
	WriteResult(data interface{}) error
}
