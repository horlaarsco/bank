module github.com/horlaarsco/bank

go 1.18

require gorm.io/gorm v1.23.4

require github.com/go-sql-driver/mysql v1.6.0 // indirect

require (
	github.com/gorilla/mux v1.8.0 // direct
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4
	gorm.io/driver/mysql v1.3.3
)
