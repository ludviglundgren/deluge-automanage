package domain

type DelugeConfig struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	Login    string `mapstructure:"login"`
	Password string `mapstructure:"password"`
}

type Rules struct {
	Enabled            bool            `mapstructure:"enabled"`
	MaxActiveDownloads int             `mapstructure:"max_active_downloads"`
	MaxDiskUsage       int             `mapstructure:"max_disk_usage"`
	MinDiskFree        int             `mapstructure:"min_disk_free"`
	Trackers           []RulesTrackers `mapstructure:"trackers"`
}

type RulesTrackers struct {
	Name        string   `mapstructure:"name"`
	Urls        []string `mapstructure:"urls"`
	MinSeedtime int      `mapstructure:"min_seed_time"`
	MinRatio    int      `mapstructure:"min_ratio"`
}

type AppConfig struct {
	Debug  bool         `mapstructure:"debug"`
	Deluge DelugeConfig `mapstructure:"deluge"`
	Rules  Rules        `mapstructure:"rules"`
}
