package ports

type DbPort interface {
	CloseDbConnection() error
	AddToHistory(answer int32, operation string) error
}
