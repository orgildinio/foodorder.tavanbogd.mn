package bootstrap

import (
	"github.com/lambda-platform/lambda/DB"
	"github.com/lambda-platform/lambda/DBSchema"
	"github.com/lambda-platform/lambda/config"
	"github.com/lambda-platform/lambda/generator"
	genertarModels "github.com/lambda-platform/lambda/generator/models"
)

func Generate() {
	var FormSchemas []genertarModels.ProjectSchemas
	var GridSchemas []genertarModels.ProjectSchemas
	var GraphqlSchemas []genertarModels.ProjectSchemas

	if config.Config.Database.Connection == "oracle" {

		DB.DB.Select("ID AS \"id\", NAME AS \"name\", SCHEMA AS \"schema\", \"TYPE\" AS \"type\", CREATED_AT AS \"created_at\", UPDATED_AT AS \"updated_at\"").Where("TYPE = ?", "form").Table("VB_SCHEMAS").Find(&FormSchemas)
		DB.DB.Select("ID AS \"id\", NAME AS \"name\", SCHEMA AS \"schema\", \"TYPE\" AS \"type\", CREATED_AT AS \"created_at\", UPDATED_AT AS \"updated_at\"").Where("TYPE = ?", "grid").Table("VB_SCHEMAS").Find(&GridSchemas)
		DB.DB.Select("ID AS \"id\", NAME AS \"name\", SCHEMA AS \"schema\", \"TYPE\" AS \"type\", CREATED_AT AS \"created_at\", UPDATED_AT AS \"updated_at\"").Where("TYPE = ?", "graphql").Table("VB_SCHEMAS").Find(&GraphqlSchemas)

	} else {
		DB.DB.Where("type = ?", "form").Table("vb_schemas").Find(&FormSchemas)
		DB.DB.Where("type = ?", "grid").Table("vb_schemas").Find(&GridSchemas)
		DB.DB.Where("type = ?", "graphql").Table("vb_schemas").Find(&GraphqlSchemas)
	}

	dbSchema := DBSchema.GetDBSchema()
	/*
	   Generate Form, Grid
	*/
	generator.ModelInit(dbSchema, FormSchemas, GridSchemas, true, false)
	/*
	   Generate GRAPHQL
	*/
	generator.GQLInit(dbSchema, GraphqlSchemas)
}
