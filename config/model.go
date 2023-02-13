package config

type Config struct {
	Server Server `yaml:"server" json:"server"`
	Mysql  Mysql  `yaml:"mysql" json:"mysql"`
	Log    Log    `yaml:"logging" json:"logging"`
	//App    App    `yaml:"app" json:"app"`
}

type Mysql struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	DB       string `yaml:"db" json:"db"`
	Username string `yaml:"user" json:"user"`
	Password string `yaml:"password" json:"password"`
}

type Server struct {
	Host string `yaml:"host" json:"host"`
	Port string `yaml:"port" json:"port"`
}

type Log struct {
	FileName string `yaml:"filename" json:"filename"`
	Levels   Levels `yaml:"level" json:"level"`
	MaxSize  int    `yaml:"maxsize" json:"maxsize"`
	MaxAge   int    `yaml:"maxage" json:"maxage"`
	Compress bool   `yaml:"compress" json:"compress"`
}

type Levels struct {
	App  string `yaml:"app" json:"app"`
	Gorm string `yaml:"gorm" json:"gorm"`
}
