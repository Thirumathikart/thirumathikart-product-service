package config

func InitConfig() {
	Environment()
	ConnectDB()
	MigrateDB()
}
