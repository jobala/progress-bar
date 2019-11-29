package core

type Config struct {
	barsize             int
	align               string
	barCompleteString   string
	barIncompleteString string
}

func GetConfig() Config {
	return Config{
		barsize:             40,
		align:               "left",
		barCompleteString:   "=",
		barIncompleteString: "-",
	}
}
