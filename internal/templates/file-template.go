package templates

func GenerateMainTemplate() string {
	mainTemplate := `
	package main

	import (
		"github.com/gorilla/mux"
		_ "github.com/jinzhu/gorm/dialects/mysql"
		"module_name/internal/routes"
		"log"
		"net/http"
	)

	func main() {
		r := mux.NewRouter()
		routes.RegisteModelAPIRoutes(r)
		http.Handle("/", r)
		log.Fatal(http.ListenAndServe("localhost:9010", r))
	}
	`
	return mainTemplate
}
func GenerateUtilsemplate() string {
	utilsTemplate := `
	package utils

	import (
		"encoding/json"
		"io"
		"net/http"
	)

	func ParseBody(r *http.Request, x interface{}) {
		if body, err := io.ReadAll(r.Body); err == nil {
			if err := json.Unmarshal([]byte(body), x); err != nil {
				return
			}
		}
	}
	`
	return utilsTemplate
}

func GenerateDBConfigTemplate() string {
	DBConfigTemplate := `
	package config

	import (
		"github.com/jinzhu/gorm"
		_ "github.com/jinzhu/gorm/dialects/mysql"
	)

	var (
		db *gorm.DB
	)

	func Connect() {
		d, err := gorm.Open("mysql", "user:password@/db-name?charset=utf8&parseTime=True&loc=Local")

		if err != nil {
			panic(err)
		}
		db = d
	}

	func GetDB() *gorm.DB {
		return db
	}
	`
	return DBConfigTemplate
}

func GenerateModelTemplate() string {
	modelTemplate := `
	package models

	import (
		"github.com/jinzhu/gorm"
		"module_name/internal/config"
	)
	
	var db *gorm.DB
	
	type Model struct {
		//GORM provides a predefined struct named gorm.Model, which includes commonly used fields like:
		//ID(uint), CreatedAt(time.Time), UpdatedAt(time.Time), DeletedAt(gorm.DeletedAt)
		//You can embed gorm.Model directly in your structs to include these fields automatically

		gorm.Model
	}
	
	func init() {
		config.Connect()
		db = config.GetDB()
		db.AutoMigrate(&Model{})
	}
	
	func (m *Model) CreateModel() *Model {
		db.NewRecord(m)
		db.Create(&m)
		return m
	}
	
	func GetAllModels() []Model {
		var models []Model
		db.Find(&models)
		return models
	}
	
	func GetModelById(id int64) (*Model, *gorm.DB) {
		var modelByID Model
		db := db.Where("ID=?", id).Find(&modelByID)
		return &modelByID, db
	}
	
	func DeleteModel(id int64) Model {
		var modelDeleted Model
		db.Where("ID=?", id).Delete(modelDeleted)
		return modelDeleted
	}
	
	`
	return modelTemplate
}

func GenerateRouteTemplate() string {
	modelTemplate := `
	package routes

	import (
		"github.com/gorilla/mux"
		"module_name/internal/controllers"
	)
	
	var RegisterModelAPIRoutes = func(r *mux.Router) {
		r.HandleFunc("/model/", controllers.CreateModel).Methods("POST")
		r.HandleFunc("/model/", controllers.GetModel).Methods("GET")
		r.HandleFunc("/model/{modelId}", controllers.GetModelById).Methods("GET")
		r.HandleFunc("/model/{modelId}", controllers.UpdateModel).Methods("PUT")
		r.HandleFunc("/model/{modelId}", controllers.DeleteModel).Methods("DELETE")
	}
	
	`
	return modelTemplate
}
