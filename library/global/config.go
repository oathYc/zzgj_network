package global

var Config *ProjectConfig

type ProjectConfig struct {
	Log      LogConfig                 `toml:"log"`    // log output path
	Server   ServiceConfig             `toml:"server"` // server log level
	Database map[string]DatabaseConfig `toml:"database"`
	// redis配置
	Redis RedisConfig `toml:"redis"`
}

// 数据库配置
type DatabaseConfig struct {
	Driver          string `toml:"driver"`            // 数据库驱动
	Source          string `toml:"source"`            // 数据库源
	ConnMaxLifeTime int    `toml:"conn_max_lifetime"` // 数据库最大连接时长
	MaxIdleConns    int    `toml:"max_id_conns"`      // 数据库最大连接数
	MaxOpenConns    int    `toml:"max_open_conns"`    // 数据库闲置连接数
}

type LogConfig struct {
	// 存储目录路径
	Dir string
	// 日志分类目录
	Category string
	// 日志级别
	Level string
	// 是否打印到控制台
	StdPrint bool `toml:"std_print"`
}

type ServiceConfig struct {
	LocalHttpAddr    string `toml:"local_http_addr"`
	MedTestAddr      string `toml:"med_test_addr"`
	MedTestStoreAddr string `toml:"med_test_store_addr"`
}

type RedisConfig struct {
	Pool map[string]RedisPoolConfig `toml:"pool"`
}

type RedisPoolConfig struct {
	Address      string `toml:"address"`
	DialPassword string `toml:"dialpassword"`
	Db           int    `toml:"db"`
	MaxIdle      int    `toml:"maxidle"`
	MaxActive    int    `toml:"maxactive"`
	Wait         bool   `toml:"wait"`
	Idletimeout  int    `toml:"idletimeout"` // 连接超时时间， 毫秒
}
