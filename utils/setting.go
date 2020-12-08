package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
AppMode  string
HttpPort  string
Db   string
DbHost  string
DbPort string
Dbuser   string
Dbname   string
Dbpwd   string
AccessKey  string
SecretKey  string
Bucket  string
QiniuServer  string
)


func init(){
	file ,err :=ini.Load("config/config.ini")
	if err !=nil{
		fmt.Println("ini.Load(../config/config.ini）faild:",err)
		return
	}
	LoadData(file)
	LoadData(file)
	Loadqiniu(file)
}
func LoadServce(file *ini.File){
	AppMode = file.Section("Server").Key("AppMode").MustString("debug")
	HttpPort =file.Section("Server").Key("HttpPort").MustString(":3000")


}
func LoadData(file *ini.File){
	Db = file.Section("Databses").Key("Db").MustString("mysql")
	DbHost  = file.Section("Databses").Key("DbHost").MustString("118.24.102.88")
	DbPort = file.Section("Databses").Key("DbPort").MustString("13306")
	Dbuser = file.Section("Databses").Key("Dbuser").MustString("root")
	Dbpwd = file.Section("Databses").Key("Dbpwd").MustString("root1234")
	Dbname = file.Section("Databses").Key("Dbname").MustString("ginblog")



}

func Loadqiniu(file *ini.File){
	AccessKey   = file.Section("Qiniu").Key("AccessKey").String()
	SecretKey   =  file.Section("Qiniu").Key("SecretKey").String()
	Bucket   = file.Section("Qiniu").Key("Bucket").String()
	QiniuServer  =file.Section("Qiniu").Key("QiniuServer").String()
}
