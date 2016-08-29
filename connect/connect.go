package connect

import (
  "log"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/SanDroid77/REST/structs"
)

var connection *gorm.DB

const engine_sql string = "mysql"
const username string = "root"
const password string = ""
const database string = "goapi"

func InitializaDatabase() {
  connection = ConnectORM(CreateString())
  log.Println("La conexion con la base de datos fue exitosa")
}

func CloseConnection() {
  connection.Close()
  log.Println("Se cerró la conexion a la base de datos")
}

func ConnectORM(stringConnection string) *gorm.DB{
  connection, err := gorm.Open(engine_sql, stringConnection)
  if err != nil {
    log.Fatal(err)
    return nil
  }
  return connection
}

func CreateString() string{
  return username + ":" + password + "@/" + database
}

func GetUser(id string) structs.User{
  user := structs.User{}
  connection.Where("id = ?", id).First(&user)
  return user
}

func CreateUser(user structs.User) structs.User{
  connection.Create(&user)
  return user
}

func UpdateUser(id string, user structs.User) structs.User{
  currentUser := structs.User{}
  connection.Where("id = ?", id).First(&currentUser)

  currentUser.Username = user.Username
  currentUser.First_Name = user.First_Name
  currentUser.Last_Name = user.Last_Name
  connection.Save(&currentUser)

  return currentUser
}

func DeleteUser(id string) {
  user := structs.User{}
  connection.Where("id = ?", id).First(&user)

  connection.Delete(&user)
}