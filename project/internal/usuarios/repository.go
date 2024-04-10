package usuarios

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error)
	LastID() (uint64, error)
}

func GetId(id int) {

}
func NewRepository() Repository {
	return &MemoryRepository{}
}
