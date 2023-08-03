package strategy

type Storage interface {
	Save(name string, data []byte) error
}

type FileStorage struct {
	strategy Storage
}

func (s *FileStorage) SetStrategy(ss Storage) {
	s.strategy = ss
}

func (s *FileStorage) Save(name string, data []byte) error {
	if s.strategy == nil {
		s.SetStrategy(new(FileSave))
	}
	return s.strategy.Save(name, data)
}
