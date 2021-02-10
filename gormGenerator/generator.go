package gormgenerator

import (
	"fmt"
	"os"
	"strings"

	//Comment m
	"github.com/dave/jennifer/jen"
	"github.com/david-yappeter/GormCrudGenerator/setting"
	"github.com/iancoleman/strcase"
)

// //PaginationVariableGenerator Pagination Variable Tools Generator
// func PaginationVariableGenerator() {
// 	err := os.Mkdir("./tools", os.ModeDir)

// 	if err == nil {

// 	} else if os.IsExist(err) {

// 	} else if os.IsNotExist(err) {

// 	} else {
// 		panic(err)
// 	}

// 	_, err = os.Open("./tools/paginationVariableGenerator.go")

// 	if err == nil {

// 	} else if os.IsExist(err) {

// 	} else if os.IsNotExist(err) {

// 	} else {
// 		panic(err)
// 	}

// 	file, err := os.Create("./tools/paginationVariableGenerator.go")

// 	if err != nil {
// 		panic(err)
// 	}

// 	//START
// 	f := jen.NewFile("tools")

// 	f.Comment("//Generated By github.com/david-yappeter/GormCrudGenerator")

// 	f.Empty()

// 	f.Comment("//PaginationVariableGenerator Pagination Variable Generator Tools")
// 	f.Func().Id("PaginationVariableGenerator").Params(jen.Id("limiter").Id("*").Int(), jen.Id("numberPage").Id("*").Int(), jen.Id("ascending").Id("*").Bool(), jen.Id("sortBy").Id("*").String(), jen.Id("filter").Index().Id("*").Int()).Params(jen.Int(), jen.Int(), jen.Bool(), jen.String(), jen.Index().Int()).
// 		Block(
// 			jen.Var().Id("limit").Int(),
// 			jen.Var().Id("page").Int(),
// 			jen.Var().Id("asc").Bool(),
// 			jen.Var().Id("sort").String(),
// 			jen.Var().Id("status").Index().Int(),

// 			jen.Empty(),

// 			jen.If(
// 				jen.Id("limiter").Op("!=").Nil(),
// 			).Block(
// 				jen.Id("limit").Op("=").Id("*limiter"),
// 			).Else().Block(
// 				jen.Id("limit").Op("=").Lit(0),
// 			),

// 			jen.Empty(),

// 			jen.If(
// 				jen.Id("numberPage").Op("!=").Nil(),
// 			).Block(
// 				jen.Id("page").Op("=").Id("*numberPage"),
// 			).Else().Block(
// 				jen.Id("page").Op("=").Lit(0),
// 			),

// 			jen.Empty(),

// 			jen.If(
// 				jen.Id("ascending").Op("!=").Nil(),
// 			).Block(
// 				jen.Id("asc").Op("=").Id("*ascending"),
// 			).Else().Block(
// 				jen.Id("asc").Op("=").True(),
// 			),

// 			jen.Empty(),

// 			jen.If(
// 				jen.Id("sortBy").Op("!=").Nil(),
// 			).Block(
// 				jen.Id("sort").Op("=").Id("*sortBy"),
// 			).Else().Block(
// 				jen.Id("sort").Op("=").Lit("id"),
// 			),

// 			jen.Empty(),

// 			jen.If(
// 				jen.Id("filter").Op("!=").Nil(),
// 			).Block(
// 				jen.For(jen.Id("_").Op(",").Id("val").Op(":=").Range().Id("filter")).Block(
// 					jen.Id("status").Op("=").Append(jen.Id("status"), jen.Id("*val")),
// 				),
// 			).Else().Block(
// 				jen.Id("status").Op("=").Index().Int().Values(),
// 			),

// 			jen.Return(jen.Id("limit"), jen.Id("page"), jen.Id("asc"), jen.Id("sort"), jen.Id("status")),
// 		)

// 	// END

// 	_, err = file.WriteString(fmt.Sprintf("%#v", f))

// 	if err != nil {
// 		panic(err)
// 	}
// }

//GormQueryToolsGenerator Pagination Query Tools Generator
func GormQueryToolsGenerator(setting setting.YamlSettings) {
	err := os.Mkdir(setting.QueryTools.Path, os.ModeDir)

	if err == nil {

	} else if os.IsExist(err) {

	} else if os.IsNotExist(err) {

	} else {
		panic(err)
	}

	_, err = os.Open(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(setting.QueryTools.Path), setting.QueryTools.Name))

	if err == nil {

	} else if os.IsExist(err) {

	} else if os.IsNotExist(err) {

	} else {
		panic(err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(setting.QueryTools.Path), setting.QueryTools.Name))

	if err != nil {
		panic(err)
	}

	//START
	f := jen.NewFile(pathGetPackage(setting.QueryTools.Path))

	f.Comment("//Generated By github.com/david-yappeter/GormCrudGenerator")

	f.Empty()

	//CREATE
	f.Comment("//QueryMaker Pagination Query Tools")
	f.Func().Id("QueryMaker").Params(
		jen.Id("query").Id("*").Qual("gorm.io/gorm", "DB"),
		jen.Id("limit").Id("*").Int(),
		jen.Id("page").Id("*").Int(),
		jen.Id("ascending").Id("*").Bool(),
		jen.Id("sortBy").Id("*").String(),
		// jen.Id("filter").Index().Id("*").Int(),
	).
		Block(

			jen.Id("sortBy").Op("=").Id("OrderBy").Call(jen.Id("sortBy"), jen.Id("ascending")),

			f.Empty(),

			// jen.Id("StatusFilter").Call(jen.Id("query"), jen.Id("filter")),

			// f.Empty(),

			jen.If(
				jen.Id("limit").Op("!=").Nil().Op("&&").Id("page").Op("!=").Nil().Op("&&").Id("*limit").Op(">").Lit(0).Op("&&").Id("*page").Op(">").Lit(0),
			).Block(
				jen.Id("offset").Op(":=").Parens(jen.Id("*page").Op("-").Lit(1)).Op("*").Id("*limit"),
				jen.Id("*query").Op("=").Id("*query").Dot("Offset").Call(jen.Id("offset")).Dot("Limit").Call(jen.Id("*limit")),
			),
			f.Empty(),

			jen.If(
				jen.Id("sortBy").Op("!=").Nil(),
			).Block(
				jen.Id("*query").Op("=").Id("*query").Dot("Order").Call(jen.Id("*sortBy")),
			),
		)

	f.Comment("//OrderBy Order By Generator (ASC,DESC)")
	f.Func().Id("OrderBy").Params(jen.Id("sortBy").Id("*").String(), jen.Id("ascending").Id("*").Bool()).Params(jen.Id("*").String()).
		Block(
			jen.If(
				jen.Id("sortBy").Op("==").Nil(),
			).Block(
				jen.Return(jen.Nil()),
			),

			jen.If(
				jen.Id("ascending").Op("!=").Nil(),
			).Block(
				jen.If(jen.Id("*ascending")).Block(
					jen.Id("sortSyntax").Op(":=").Id("*sortBy").Op("+").Lit(" ASC"),
					jen.Return(jen.Id("&sortSyntax")),
				),
				jen.Id("sortSyntax").Op(":=").Id("*sortBy").Op("+").Lit(" DESC"),
				jen.Return(jen.Id("&sortSyntax")),
			),

			jen.Return(jen.Nil()),
		)

	// f.Comment("//StatusFilter Filtering By Status")
	// f.Func().Id("StatusFilter").Params(jen.Id("query").Id("*").Qual("gorm.io/gorm", "DB"), jen.Id("filter").Index().Int()).
	// 	Block(

	// 		jen.Id("filterLen").Op(":=").Len(jen.Id("filter")),

	// 		jen.Empty(),

	// 		jen.If(
	// 			jen.Id("filterLen").Op(">").Lit(0),
	// 		).Block(
	// 			jen.Id("*query").Op("=").Id("*query").Dot("Where").Call(jen.Lit("status IN (?)"), jen.Id("filter")),
	// 		),
	// 	)

	// END

	_, err = file.WriteString(fmt.Sprintf("%#v", f))

	if err != nil {
		panic(err)
	}
}

//GormLogGenerator Gorm Log Setting
func GormLogGenerator(setting setting.YamlSettings) {
	err := os.Mkdir(setting.Database.Setting.Path, os.ModeDir)

	if err == nil {

	} else if os.IsExist(err) {

	} else if os.IsNotExist(err) {

	} else {
		panic(err)
	}

	_, err = os.Open(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(setting.Database.Setting.Path), setting.Database.Setting.Name))

	if err == nil {

	} else if os.IsExist(err) {

	} else if os.IsNotExist(err) {

	} else {
		panic(err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(setting.Database.Setting.Path), setting.Database.Setting.Name))

	if err != nil {
		panic(err)
	}

	//START
	f := jen.NewFile(pathGetPackage(setting.Database.Setting.Path))

	f.Comment("//Generated By github.com/david-yappeter/GormCrudGenerator")

	f.Empty()

	//CREATE

	f.Comment("//InitConfig Initialize Config")
	f.Func().Id("InitConfig").Params().Params(jen.Id("*").Qual("gorm.io/gorm", "Config")).
		Block(
			jen.Return(jen.Id("&").Qual("gorm.io/gorm", "Config").Values(
				jen.DictFunc(func(d jen.Dict) {
					d[jen.Id("Logger")] = jen.Id("InitLog").Call()
					d[jen.Id("NamingStrategy")] = jen.Id("InitNamingStrategy").Call()
				},
				),
			)),
		)

	f.Comment("//InitLog Connection Log Configuration")
	f.Func().Id("InitLog").Params().Params(jen.Qual("gorm.io/gorm/logger", "Interface")).
		Block(

			jen.Id("newLogger").Op(":=").Qual("gorm.io/gorm/logger", "New").Call(
				jen.Qual("log", "New").Call(
					jen.Qual("os", "Stdout"),
					jen.Lit("\r\n"),
					jen.Qual("log", "LstdFlags"),
				),

				// jen.Empty(),

				jen.Qual("gorm.io/gorm/logger", "Config").Values(
					jen.DictFunc(func(d jen.Dict) {
						if setting.Database.Setting.SlowThreshold > 1 {
							d[jen.Id("SlowThreshold")] = jen.Qual("time", "Second").Op("*").Lit(setting.Database.Setting.SlowThreshold)
						} else {
							d[jen.Id("SlowThreshold")] = jen.Qual("time", "Second")
						}
						d[jen.Id("LogLevel")] = jen.Qual("gorm.io/gorm/logger", setting.Database.Setting.LogLevel[0])
						d[jen.Id("Colorful")] = jen.True()
					},
					),
				),
			),

			jen.Return(jen.Id("newLogger")),
		)

	f.Comment("//InitNamingStrategy Init NamingStrategy")
	f.Func().Id("InitNamingStrategy").Params().Params(jen.Id("*").Qual("gorm.io/gorm/schema", "NamingStrategy")).
		Block(
			jen.Return(
				jen.Id("&").Qual("gorm.io/gorm/schema", "NamingStrategy").Values(
					jen.DictFunc(func(d jen.Dict) {
						d[jen.Id("TablePrefix")] = jen.Lit(setting.Database.Setting.TablePrefix)
						d[jen.Id("SingularTable")] = jen.Lit(setting.Database.Setting.SingularTable)
					},
					),
				),
			),
		)

	// END

	_, err = file.WriteString(fmt.Sprintf("%#v", f))

	if err != nil {
		panic(err)
	}
}

//GormConnectionGenerator Gorm Connection Gorm V2
func GormConnectionGenerator(setting setting.YamlSettings, goModName string, connectionType string) {

	err := os.Mkdir(setting.Database.Path, os.ModeDir)

	if err == nil {

	} else if os.IsExist(err) {

	} else if os.IsNotExist(err) {

	} else {
		panic(err)
	}

	_, err = os.Open(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(setting.Database.Path), setting.Database.Name))

	if err == nil {

	} else if os.IsExist(err) {

	} else if os.IsNotExist(err) {

	} else {
		panic(err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.go", pathRemoveLastStrip(setting.Database.Path), setting.Database.Name))

	if err != nil {
		panic(err)
	}

	//START
	f := jen.NewFile(pathGetPackage(setting.Database.Path))

	f.Comment("//Generated By github.com/david-yappeter/GormCrudGenerator")

	f.Empty()

	//CREATE
	if connectionType == "mysql" {
		f.Comment("//ConnectGorm Database Connection to Gorm V2")
		f.Func().Id("ConnectGorm").Params().Params(jen.Id("*").Qual("gorm.io/gorm", "DB")).
			Block(
				jen.Id("databaseConfig").Op(":=").Qual("fmt", "Sprintf").Call(
					jen.Lit("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true"),
					jen.Qual("os", "Getenv").Call(jen.Lit("DB_USER")),
					jen.Qual("os", "Getenv").Call(jen.Lit("DB_PASSWORD")),
					jen.Qual("os", "Getenv").Call(jen.Lit("DB_HOST")),
					jen.Qual("os", "Getenv").Call(jen.Lit("DB_PORT")),
					jen.Qual("os", "Getenv").Call(jen.Lit("DB_DATABASE")),
				),

				jen.Empty(),

				jen.Id("db").Op(",").Err().Op(":=").Qual("gorm.io/gorm", "Open").Params(jen.Qual("gorm.io/driver/mysql", "Open").Params(jen.Id("databaseConfig")), jen.Qual(goModName+"/logger", "InitConfig").Call()),

				jen.Empty(),

				jen.If(
					jen.Err().Op("!=").Nil(),
				).Block(
					jen.Qual("fmt", "Println").Call(jen.Err()),
					jen.Panic(jen.Lit("Fail To Connect Database")),
				),

				jen.Empty(),

				jen.Return(jen.Id("db")),
			)
	} else if connectionType == "postgres" {
		f.Comment("//ConnectGorm Database Connection to Gorm V2")
		f.Func().Id("ConnectGorm").Params().Params(jen.Id("*").Qual("gorm.io/gorm", "DB")).
			Block(
				// jen.Id("databaseConfig").Op(":=").Qual("fmt", "Sprintf").Call(
				// 	jen.Lit("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true"),
				// 	jen.Qual("os", "Getenv").Call(jen.Lit("DB_USER")),
				// 	jen.Qual("os", "Getenv").Call(jen.Lit("DB_PASSWORD")),
				// 	jen.Qual("os", "Getenv").Call(jen.Lit("DB_HOST")),
				// 	jen.Qual("os", "Getenv").Call(jen.Lit("DB_PORT")),
				// 	jen.Qual("os", "Getenv").Call(jen.Lit("DB_DATABASE")),
				// ),

				jen.Id("dsn").Op(":=").Qual("os", "Getenv").Call(jen.Lit("POSTGRE_URI")),

				jen.Empty(),

				jen.Id("db").Op(",").Err().Op(":=").Qual("gorm.io/gorm", "Open").Params(jen.Qual("gorm.io/driver/postgres", "New").Params(jen.Qual("gorm.io/driver/postgres", "Config").Values(
					jen.DictFunc(func(d jen.Dict) {
						d[jen.Id("DSN")] = jen.Id("dsn")
						d[jen.Id("PreferSimpleProtocol")] = jen.True()
					}),
				)), jen.Qual(goModName+"/logger", "InitConfig").Call()),

				jen.Empty(),

				jen.If(
					jen.Err().Op("!=").Nil(),
				).Block(
					jen.Qual("fmt", "Println").Call(jen.Err()),
					jen.Panic(jen.Lit("Fail To Connect Database")),
				),

				jen.Empty(),

				jen.Return(jen.Id("db")),
			)
	}

	// END

	_, err = file.WriteString(fmt.Sprintf("%#v", f))

	if err != nil {
		panic(err)
	}

}

//CrudGenerator Crud Generator
func CrudGenerator(setting setting.YamlSettings, nameGoMod string, nameStruct string, attribute map[string][]string) error {

	structName := nameStruct
	goModName := nameGoMod

	if !checkIgnore(setting.Service.From.Ignore, structName) {
		return nil
	}

	camelStructName := strcase.ToLowerCamel(structName)
	snakeStructName := strcase.ToSnake(structName)

	attributeMap := make(map[string][]string)
	attributeMap = attribute

	f := jen.NewFile(pathGetPackage(setting.Service.To.Path))

	f.Comment("//Generated By github.com/david-yappeter/GormCrudGenerator")

	f.Empty()

	//CREATE
	f.Comment("//" + structName + "Create Create")
	f.Func().Id(structName+"Create").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("input").Qual(goModName+"/graph/model", "New"+structName),
	).Params(jen.Id("*").Qual(goModName+"/graph/model", structName), jen.Error()).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Id(camelStructName).Op(":=").Qual(goModName+"/graph/model", structName).Values(
				jen.DictFunc(func(d jen.Dict) {
					for _, val := range attributeMap[structName] {
						if strings.Trim(val, "\t") == "ID" {
							continue
						}
						d[jen.Id(val)] = jen.Id("input").Dot(val)
					}
				}),
			),

			jen.Empty(),

			jen.Id("err").Op(":=").Id("db").Dot("Table").Params(jen.Lit(snakeStructName)).Dot("Create").Params(jen.Id("&"+camelStructName)).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Id("&"+camelStructName), jen.Nil()),
		)

	//CREATE BATCH
	f.Comment("//" + structName + "CreateBatch Create Batch")
	f.Func().Id(structName+"CreateBatch").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("input").Id("[]*").Qual(goModName+"/graph/model", "New"+structName),
	).Params(jen.Id("[]*").Qual(goModName+"/graph/model", structName), jen.Error()).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Var().Id(camelStructName+"Batch").Id("[]*").Qual(goModName+"/graph/model", structName),

			jen.Empty(),

			jen.For(
				jen.Id("_,").Id("val").Op(":=").Range().Id("input"),
			).Block(
				jen.Id(camelStructName).Op(":=").Qual(goModName+"/graph/model", structName).Values(
					jen.DictFunc(func(d jen.Dict) {
						for _, val := range attributeMap[structName] {
							if strings.Trim(val, "\t") == "ID" {
								continue
							}
							d[jen.Id(val)] = jen.Id("val").Dot(val)
						}
					}),
				),
				jen.Empty(),

				jen.Id(camelStructName+"Batch").Op("=").Append(jen.Id(camelStructName+"Batch"), jen.Id("&"+camelStructName)),
			),

			jen.Empty(),

			jen.Id("err").Op(":=").Id("db").Dot("Table").Params(jen.Lit(snakeStructName)).Dot("Create").Params(jen.Id("&"+camelStructName+"Batch")).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Id(camelStructName+"Batch"), jen.Nil()),
		)

	//UPDATE
	f.Comment("//" + structName + "Update Update")
	f.Func().Id(structName+"Update").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("input").Qual(goModName+"/graph/model", "Update"+structName),
	).Params(jen.Id("*").Qual(goModName+"/graph/model", structName), jen.Error()).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Id(camelStructName).Op(":=").Qual(goModName+"/graph/model", structName).Values(
				jen.DictFunc(func(d jen.Dict) {
					for _, val := range attributeMap[structName] {
						d[jen.Id(val)] = jen.Id("input").Dot(val)
					}
				}),
			),

			jen.Empty(),

			jen.Id("err").Op(":=").Id("db").Dot("Table").Call(jen.Lit(snakeStructName)).Dot("Where").Call(jen.Lit("id = ?"), jen.Id("input.ID")).Dot("Updates").Call(
				jen.Map(jen.String()).Interface().Block(
					jen.DictFunc(
						func(d jen.Dict) {
							for _, val := range attributeMap[structName] {
								d[jen.Lit(strcase.ToSnake(val))] = jen.Id("input").Dot(val)
							}
						},
					),
				),
			).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Id("&"+camelStructName), jen.Nil()),
		)

	//DELETE
	f.Comment("//" + structName + "Delete Delete")
	f.Func().Id(structName+"Delete").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("id").Int(),
	).Params(jen.String(), jen.Error()).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Id("err").Op(":=").Id("db").Dot("Table").Params(jen.Lit(snakeStructName)).Dot("Where").Params(jen.Lit("id = ?"), jen.Id("id")).Dot("Delete").Params(jen.Id("&model."+structName+"{}")).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Lit("Fail"), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Lit("Success"), jen.Nil()),
		)

	//GET BY ID
	f.Comment("//" + structName + "GetByID Get By ID")
	f.Func().Id(structName+"GetByID").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("id").Int(),
	).Params(jen.Id("*").Qual(goModName+"/graph/model", structName), jen.Error()).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Var().Id(camelStructName).Qual(goModName+"/graph/model", structName),

			jen.Empty(),

			jen.Id("err").Op(":=").Id("db").Dot("Table").Params(jen.Lit(snakeStructName)).Dot("Where").Params(jen.Lit("id = ?"), jen.Id("id")).Dot("First").Params(jen.Id("&"+camelStructName)).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Id("&"+camelStructName), jen.Nil()),
		)

	//GET ALL
	f.Comment("//" + structName + "GetAll GetAll")
	f.Func().Id(structName+"GetAll").Params(
		jen.Id("ctx").Qual("context", "Context"),
	).Params(jen.Id("[]*").Qual(goModName+"/graph/model", structName), jen.Error()).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Var().Id(camelStructName).Id("[]*").Qual(goModName+"/graph/model", structName),

			jen.Empty(),

			jen.Id("err").Op(":=").Id("db").Dot("Table").Params(jen.Lit(snakeStructName)).Dot("Find").Params(jen.Id("&"+camelStructName)).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Id(camelStructName), jen.Nil()),
		)

	//PAGINATION
	f.Comment("//" + structName + "Pagination Pagination")
	f.Func().Id(structName+"Pagination").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("*limit").Int(),
		jen.Id("*page").Int(),
		jen.Id("*ascending").Bool(),
		jen.Id("*sortBy").String(),
		// jen.Id("filter").Index().Int(),
	).Params(
		jen.Id("[]*").Qual(goModName+"/graph/model", structName),
		jen.Error(),
	).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Var().Id(camelStructName).Id("[]*").Qual(goModName+"/graph/model", structName),

			jen.Empty(),

			jen.Id("query").Op(":=").Id("db").Dot("Table").Call(jen.Lit(snakeStructName)),

			jen.Empty(),

			jen.Qual(goModName+"/tools", "QueryMaker").Call(
				jen.Id("query"),
				jen.Id("limit"),
				jen.Id("page"),
				jen.Id("ascending"),
				jen.Id("sortBy"),
				// jen.Id("filter"),
			),

			jen.Err().Op(":=").Id("query").Dot("Find").Call(jen.Id("&"+camelStructName)).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Nil(), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Id(camelStructName), jen.Nil()),
		)

	//TOTAL DATA PAGINATION
	f.Comment("//" + structName + "TotalDataPagination  Total Data Pagination")
	f.Func().Id(structName+"TotalDataPagination ").Params(
		jen.Id("ctx").Qual("context", "Context"),
		jen.Id("*limit").Int(),
		jen.Id("*page").Int(),
		jen.Id("*ascending").Bool(),
		jen.Id("*sortBy").String(),
	).Params(
		jen.Int(),
		jen.Error(),
	).
		Block(
			jen.Id("db").Op(":=").Qual(goModName+"/config", "ConnectGorm").Call(),
			jen.Id("sqlDB").Op(",").Id("_").Op(":=").Id("db").Dot("DB").Call(),
			jen.Defer().Id("sqlDB").Dot("Close").Call(),

			jen.Empty(),

			jen.Var().Id("count").Int64(),

			jen.Empty(),

			jen.Id("query").Op(":=").Id("db").Dot("Table").Call(jen.Lit(snakeStructName)),

			jen.Empty(),

			jen.Qual(goModName+"/tools", "QueryMaker").Call(
				jen.Id("query"),
				jen.Id("limit"),
				jen.Id("page"),
				jen.Id("ascending"),
				jen.Id("sortBy"),
			),

			jen.Err().Op(":=").Id("query").Dot("Count").Call(jen.Id("&count")).Dot("Error"),

			jen.Empty(),

			jen.If(
				jen.Id("err").Op("!=").Nil(),
			).Block(
				jen.Qual("fmt", "Println").Call(jen.Id("err")),
				jen.Return(jen.Lit(0), jen.Id("err")),
			),

			jen.Empty(),

			jen.Return(jen.Int().Parens(jen.Id("count")), jen.Nil()),
		)

	err := os.Mkdir(setting.Service.To.Path, os.ModeDir)

	if err == nil {

	} else if os.IsExist(err) {

	} else if os.IsNotExist(err) {

	} else {
		panic(err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s%s.go", setting.Service.To.Path, camelStructName, setting.Service.To.Postfix))

	if err != nil {
		return err
	}

	_, err = file.WriteString(fmt.Sprintf("%#v", f))

	if err != nil {
		return err
	}

	return nil

}

func pathRemoveLastStrip(path string) string {
	if len(path) > 0 && path[len(path)-1] == []byte("/")[0] {
		return path[0 : len(path)-1]
	}
	return path
}

func pathGetPackage(path string) string {
	splitted := strings.Split(path, "/")
	return splitted[len(splitted)-1]
}

//Case Sensitive, true -> pass, false -> ignore it (skip)
func checkIgnore(ignore []string, structName string) bool {
	for _, val := range ignore {
		if strings.Trim(val, " ") == strings.Trim(structName, " ") {

			return false
		}
	}
	return true
}
