package newsreader

import (
	. "Gideon/internal/config"
)

func InitNewsReader() *NewsReader {
	return &NewsReader{
		ApiKey: AppConfig.NewsReader.APIKey,
	}
}
