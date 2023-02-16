package config

type Config struct {
	Server  Server  `yaml:"server" json:"server"`
	Mysql   Mysql   `yaml:"mysql" json:"mysql"`
	Logging Logging `yaml:"logging" json:"logging"`
	//App    App    `yaml:"app" json:"app"`
}

type Mysql struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	DB       string `yaml:"db" json:"db"`
	User     string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
}

type Server struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type Logging struct {
	FileName string `yaml:"filename" json:"filename"`
	Level    Level  `yaml:"level" json:"level"`
	MaxSize  int    `yaml:"maxsize" json:"maxsize"`
	MaxAge   int    `yaml:"maxage" json:"maxage"`
	Compress bool   `yaml:"compress" json:"compress"`
}

type Level struct {
	App  string `yaml:"app" json:"app"`
	Gorm string `yaml:"gorm" json:"gorm"`
}
