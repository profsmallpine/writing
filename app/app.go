package app

import (
	"embed"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/profsmallpine/writing/http/web"
	"github.com/profsmallpine/writing/migrations"
	"github.com/xy-planning-network/trails/http/template"
	"github.com/xy-planning-network/trails/postgres"
	"github.com/xy-planning-network/trails/ranger"
)

func New(files embed.FS) (*ranger.Ranger, error) {
	_ = godotenv.Load()

	config := &postgres.CxnConfig{IsTestDB: false, URL: os.Getenv("DATABASE_URL")}
	if config.URL == "" {
		config.Host = envVarOrString("PG_HOST", "localhost")
		config.Name = envVarOrString("PG_NAME", "mid_dev_db")
		config.Password = envVarOrString("PG_PASSWORD", "")
		config.Port = envVarOrString("PG_PORT", "5432")
		config.User = envVarOrString("PG_USER", "mid_dev_db_user")
	}
	db, err := postgres.Connect(config, migrations.List)
	if err != nil {
		return nil, err
	}

	rng, err := ranger.New(
		ranger.WithDB(postgres.NewService(db)),
		ranger.WithUserSessions(userStorer{}),
		ranger.DefaultParser(
			files,
			template.WithFn("add1", func(value int) int { return value + 1 }),
			template.WithFn("isLast", func(idx, length int) bool { return (idx + 1) == length }),
			template.WithFn("minus1", func(value int) int { return value - 1 }),
		),
	)
	if err != nil {
		return nil, err
	}

	handler := web.NewHandler(
		db,
		rng,
		envVarOrString("WRITER_KEY", "super-secret"),
		strings.Split(envVarOrString("WHITELIST_IPS", "0.0.0.0"), ","),
	)
	handler.SetupRoutes()

	return rng, nil
}

func envVarOrString(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
