package models

type Flags struct {
	Configured bool          `long:"configured" env:"CONFIGURED" required:"true" description:"If set to false, the application will exit, should be set to true when everything is configured correctly"`
	Env        string        `short:"e" long:"env" env:"ENV" required:"true" description:"The environment the program is running in: production/development"`
	DB         ConfigDB      `group:"Database Options" namespace:"db" env-namespace:"DB"`
	Redis      ConfigRedis   `group:"Redis Options" namespace:"redis" env-namespace:"REDIS"`
	Discord    ConfigDiscord `group:"Discord Options" namespace:"bot" env-namespace:"BOT"`
}

type ConfigDB struct {
	URL string `env:"URL" long:"url" required:"true" description:"database connection url"`
}

type ConfigRedis struct {
	Addr     string `long:"addr" env:"ADDR" description:"Redis server address" required:"true"`
	Password string `long:"password" env:"PASSWORD" description:"Redis server password" required:"true"`
	DB       int    `long:"db" env:"DB" description:"Redis server database" required:"true"`
}

type ConfigDiscord struct {
	Token        string `long:"token" env:"TOKEN" description:"Discord Bot Token" required:"true"`
	OwnerID      string `long:"owner" env:"OWNER_ID" description:"Owner of the discord bot" required:"true"`
	ErrorServer  string `long:"error-server-id" env:"ERROR_SERVER_ID" description:"Server ID for the error log server" required:"true"`
	ErrorChannel string `long:"error-channel-id" env:"ERROR_CHANNEL_ID" description:"Channel ID for the error log channel" required:"true"`
}
