package usuarios

var Usuarios []Usuario
var lastID uint64 = 0

type MemoryRepository struct {
}

func (m *MemoryRepository) GetAll() ([]Usuario, error) {
	return Usuarios, nil
}

func (m *MemoryRepository) GetId() (Usuario, error) {
	user := Usuario{}
	return user, nil
}

func (m *MemoryRepository) Store(nome string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {
	lastID++
	p := Usuario{
		Id:          lastID,
		Nome:        nome,
		Sobrenome:   sobrenome,
		Email:       email,
		Idade:       idade,
		Altura:      altura,
		Ativo:       ativo,
		DataCriacao: datacriacao,
	}
	Usuarios = append(Usuarios, p)
	return p, nil
}

func (m *MemoryRepository) LastID() (uint64, error) {
	return lastID, nil
}
