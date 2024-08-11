package database

import (
	"fmt"
	"log"
	env "mystic-forge-api/src/shared"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var port string = env.GetEnv("db_port")
var user string = env.GetEnv("db_user")
var host string = env.GetEnv("db_host")
var password string = env.GetEnv("db_password")
var name string = env.GetEnv("db_name")

var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	user,
	password,
	host,
	port,
	name,
)

func FactorieMysql() gorm.DB {
	var mysqlDsn = mysql.Open(dsn)
	var connection = DatabaseInit(mysqlDsn)

	if connection.Err != nil {
		fmt.Println("Erro ao pingar o banco de dados:", connection.Err)
		log.Fatalln(connection.Err)
	}

	fmt.Println("Conexão com o banco de dados bem-sucedida")
	return *connection.database
}
