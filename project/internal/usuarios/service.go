package usuarios

type Service interface {
	GetAll() ([]Usuario, error)
	Store(name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Usuario, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (s *service) GetId(id int) (Usuario, error) {
	user, err := s.repository.GetId(id)
	if err != nil {
		return Usuario{}, err
	}

	return user, nil

}

func (s *service) Store(name string, sobrenome string, email string, idade int, altura int, ativo bool, datacriacao string) (Usuario, error) {

	usuarios, err := s.repository.Store(name, sobrenome, email, idade, altura, ativo, datacriacao)
	if err != nil {
		return Usuario{}, err
	}

	return usuarios, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
