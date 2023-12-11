package main

import (
	"github.com/peifengll/tests/testgorm/testgen/gendemo/dal/model"
	"github.com/peifengll/tests/testgorm/testgen/gendemo/dal/repository"
	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(model.Floder{})

	g.ApplyInterface(func(repository.FloderRepositoryG) {}, model.Floder{})

	g.Execute()
}
