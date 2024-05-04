package templates

import "fmt"

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
		routes.RegisterModelAPIRoutes(r)
		http.Handle("/", r)
		log.Fatal(http.ListenAndServe("localhost:9010", r))
	}
	`
	return mainTemplate
}
func GenerateUtilsTemplate() string {
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
func GenerateControllerTemplate() string {
	controllerTemplate := `
	package controllers

	import (
		"encoding/json"
		"fmt"
		"github.com/gorilla/mux"
		"module_name/internal/models"
		"module_name/internal/utils"
		"net/http"
		"strconv"
	)


	func GetModels(w http.ResponseWriter, r *http.Request) {
		newModels := models.GetAllModels()
		res, _ := json.Marshal(newModels)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	func GetModelById(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ModelId := vars["modelId"]
		ID, err := strconv.ParseInt(ModelId, 0, 0)

		if err != nil {
			fmt.Println("Error while parsing")
		}
		ModelDetails, _ := models.GetModelById(ID)

		res, _ := json.Marshal(ModelDetails)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	func CreateModel(w http.ResponseWriter, r *http.Request) {
		CreatedModel := &models.Model{}
		utils.ParseBody(r, CreatedModel)

		m := CreatedModel.CreateModel()

		res, _ := json.Marshal(m)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	func DeleteModel(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ModelId := vars["modelId"]
		ID, err := strconv.ParseInt(ModelId, 0, 0)

		if err != nil {
			fmt.Println("Error while parsing")
		}

		deletedModel := models.DeleteModel(ID)

		res, _ := json.Marshal(deletedModel)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	func UpdateModel(w http.ResponseWriter, r *http.Request) {
		var UpdateModel = &models.Model{}
		utils.ParseBody(r, UpdateModel)
		vars := mux.Vars(r)
		ModelId := vars["modelId"]
		ID, err := strconv.ParseInt(ModelId, 0, 0)
		if err != nil {
			fmt.Println("Error while parsing")
		}
		ModelDetails, db := models.GetModelById(ID)
		if UpdateModel.Name != "" {
			ModelDetails.Name = UpdateModel.Name
		}
		db.Save(&ModelDetails)
		res, _ := json.Marshal(ModelDetails)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	`
	return controllerTemplate
}

func GenerateDBConfigDefaultTemplate() string {
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

func GenerateDBConfigTemplate(dbConfig AdolfDBConfig) string {
	switch dbConfig.DBName {
	case POSTGRES.String():
		DBConfigTemplate := `
		package config
	
		import (
			"gorm.io/driver/postgres"
			"gorm.io/gorm"
		
			"log"
		)
		
		var (
			db *gorm.DB
		)
		
		func Connect() {
			dsn := "host=%s user=%s password=%s dbname=%s port=%d"
		
			d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatal("Database error: ", err)
			}
		
			db = d
		}
		
		func GetDB() *gorm.DB {
			return db
		}	
		`
		formatted := fmt.Sprintf(string(DBConfigTemplate), dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)

		return formatted
	case MYSQL.String():
		DBConfigTemplate := `
		package config
	
		import (
			"gorm.io/driver/mysql"
			"gorm.io/gorm"
		
			"log"
		)
		
		var (
			db *gorm.DB
		)
		
		func Connect() {
			dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		
			d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatal("Database error: ", err)
			}
		
			db = d
		}
		
		func GetDB() *gorm.DB {
			return db
		}	
		`
		formatted := fmt.Sprintf(string(DBConfigTemplate), dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

		return formatted
	case SQLITE.String():
		DBConfigTemplate := `
		package config
	
		import (
			"gorm.io/driver/sqlite" // Sqlite driver based on CGO
			// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
			"gorm.io/gorm"
		
			"log"
		)
		
		var (
			db *gorm.DB
		)
		
		func Connect() {
			d, err := gorm.Open(sqlite.Open("%s.db"), &gorm.Config{})
			if err != nil {
				log.Fatal("Database error: ", err)
			}
		
			db = d
		}
		
		func GetDB() *gorm.DB {
			return db
		}	
		`
		formatted := fmt.Sprintf(string(DBConfigTemplate), dbConfig.DBName)

		return formatted
	default:
		return GenerateDBConfigDefaultTemplate()
	}

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
		Name        string
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
		r.HandleFunc("/model/", controllers.GetModels).Methods("GET")
		r.HandleFunc("/model/{modelId}", controllers.GetModelById).Methods("GET")
		r.HandleFunc("/model/{modelId}", controllers.UpdateModel).Methods("PUT")
		r.HandleFunc("/model/{modelId}", controllers.DeleteModel).Methods("DELETE")
	}
	
	`
	return modelTemplate
}
