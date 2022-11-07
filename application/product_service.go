package application

type ProductService struct {
	Persistence IProductPersistenceInterface
}

func (s *ProductService) Get(id string) (IProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
