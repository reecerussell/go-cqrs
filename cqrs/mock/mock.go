//go:generate mockgen -package=mock -source=../command.go -destination=command.go
//go:generate mockgen -package=mock -source=../query.go -destination=query.go
//go:generate mockgen -package=mock -source=../command_executor.go -destination=command_executor.go
//go:generate mockgen -package=mock -source=../query_executor.go -destination=query_executor.go

package mock
