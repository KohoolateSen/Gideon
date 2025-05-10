package bot

import (
	"errors"
	"fmt"
	"strings"
)

func (tb *TelegramBot) CommandArgValidator(command string, args_len int, min int, max int) error {
	if args_len < min || args_len > max {
		return fmt.Errorf("incorrect number of arguments for %s", command)
	}

	return nil
}

func (tb *TelegramBot) InlineQueryArgValidator(query string) ([]string, error) {
	query_args := strings.Split(query, "@")

	if len(query_args) == 2 {
		return query_args, nil
	}

	return nil, errors.New("")
}
