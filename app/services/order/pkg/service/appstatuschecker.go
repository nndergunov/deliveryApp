package service

type AppStatusChecker interface {
	CheckStatuses() error
}
