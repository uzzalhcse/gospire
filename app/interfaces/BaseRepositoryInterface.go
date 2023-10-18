package interfaces

type BaseRepositoryInterface interface {
	GetAll(dest interface{}) error
	GetByID(dest interface{}, id string) error
	GetWhere(dest interface{}, condition string, args ...interface{}) error
	GetFirst(dest interface{}, condition string, args ...interface{}) error
	Create(data interface{}) error
	Update(data interface{}) error
	Delete(data interface{}) error
}
