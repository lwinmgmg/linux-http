package env

import (
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	LH_HOST    string
	LH_PORT    int
	LH_ISSUERS []string
	LH_SECRET  string
	LH_DB_PATH string
}

func getEnvString(key, val string) string {
	out, ok := os.LookupEnv(key)
	if ok {
		return out
	}
	return val
}

func getEnvInt(key string, val int) int {
	out, ok := os.LookupEnv(key)
	if ok {
		if out1, err := strconv.Atoi(out); err == nil {
			return out1
		}
	}
	return val
}

func NewEnv() Settings {
	lhSecret, ok := os.LookupEnv("LH_SECRET")
	if !ok {
		panic("Can't start without LH_SECRET")
	}
	return Settings{
		LH_HOST:    getEnvString("LH_HOST", "localhost"),
		LH_PORT:    getEnvInt("LH_PORT", 80),
		LH_ISSUERS: strings.Split(getEnvString("LH_ISSUERS", "linux-http"), ","),
		LH_SECRET:  lhSecret,
		LH_DB_PATH: getEnvString("LH_DB_PATH", "db/lh.db"),
	}
}
