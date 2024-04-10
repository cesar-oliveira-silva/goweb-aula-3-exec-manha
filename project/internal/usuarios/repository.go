package usuarios

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error)
	LastID() (uint64, error)
	Update(Id uint64, name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error)
	UpdateName(id uint64, name string) (Usuario, error)
	Delete(id uint64) error
}

func NewRepository() Repository {
	return &MemoryRepository{}
}
