package structs

type Config struct {
	Host   string `yaml:"host"`
	DbName string `yaml:"dbName"`
	DbAddr string `yaml:"dbAddr"`
}

type SecretConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
