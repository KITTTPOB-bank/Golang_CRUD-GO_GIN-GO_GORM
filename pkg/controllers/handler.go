package controllers

import (
	"backenddemo/pkg/dbconfig"
	"backenddemo/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ตั้งค่าและเรียกใช้ ฐานข้อมูล
var db *gorm.DB

func init() {
	dbconfig.Connect()
	db = dbconfig.GetDB()
	fmt.Println("Successfully connected to database!", db)
}

type Characters models.CharacterData

func GetallCharacter(c *gin.Context) {
	var getalldata []Characters
	db.Table("Characters").Find(&getalldata)
	json.NewEncoder(c.Writer).Encode(getalldata)

}
func CreateCharacter(c *gin.Context) {
	db.AutoMigrate(&Characters{})
	var getdatabyid = &Characters{}
	if err := json.NewDecoder(c.Request.Body).Decode(&getdatabyid); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	Newcharacter := Characters{
		Character: getdatabyid.Character,
	}

	if err := db.Table("Characters").Create(&Newcharacter).Error; err != nil {
		return
	}
	json.NewEncoder(c.Writer).Encode(&Newcharacter)
}

func GetcharacterbyID(c *gin.Context) {
	var getdatabyid Characters

	getid := c.Param("id")
	Id, _ := strconv.Atoi(getid)
	db.Table("Characters").Where("ID=?", Id).Find(&getdatabyid)
	json.NewEncoder(c.Writer).Encode(&getdatabyid)
}

func UpdateCharacter(c *gin.Context) {
	var updateData = &Characters{}
	var getdatabyid Characters
	if err := json.NewDecoder(c.Request.Body).Decode(&updateData); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	databyid := db.Table("Characters").Where("ID=?", updateData.Id).Find(&getdatabyid)
	if updateData.Character != "" {
		getdatabyid.Character = updateData.Character
	}

	databyid.Save(&getdatabyid)

	json.NewEncoder(c.Writer).Encode(&getdatabyid)
}

func DeleteCharacterbyID(c *gin.Context) {
	var data Characters
	var dropData = &Characters{}
	if err := json.NewDecoder(c.Request.Body).Decode(&dropData); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	db.Table("Characters").Where("ID=?", dropData.Id).Delete(data)
}
