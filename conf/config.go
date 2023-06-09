package conf

import (
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DbPathRead  string
	DbPathWrite string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	Host        string
	ProductPath string
	AvatarPath  string

	ValidatorLang string

	LogSavePath string
	LogFileName string
	LogFileExt  string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Print("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadMysql(file)
	LoadEmail(file)
	LoadPhotoPath(file)
	LoadValidator(file)
	DbPathRead = strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	DbPathWrite = strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").String()
	ProductPath = file.Section("path").Key("ProductPath").String()
	AvatarPath = file.Section("path").Key("AvatarPath").String()
}

func LoadValidator(file *ini.File) {
	ValidatorLang = file.Section("validator").Key("Language").String()
}

func LoadLogger(file *ini.File) {
	LogSavePath = file.Section("logger").Key("LogSavePath").String()
	LogFileName = file.Section("logger").Key("LogFileName").String()
	LogFileExt = file.Section("logger").Key("LogFileExt").String()
}
