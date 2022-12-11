package providers

import (
	dummydata "regin/service/dummy_data"

	"github.com/sarulabs/di"
)

func Init() (*di.Container, error) {
	builder, err := di.NewBuilder()

	builder.Add(di.Def{
		Name: "dummy-data",
		Build: func(cnt di.Container) (interface{}, error) {
			return dummydata.Default()
		},
	})

	cnt := builder.Build()

	return &cnt, err
}
