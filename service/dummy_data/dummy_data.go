package dummydata

type DummyDataService struct {
}

func (dd *DummyDataService) GetData() string {
	return "Dummy-Data"
}

func Default() (interface{}, error) {
	return &DummyDataService{}, nil
}
