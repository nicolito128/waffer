package flags

import (
	"fmt"
	"strings"
)

const TokenPrefix string = "--"

type Token struct {
	Name        string
	Optional    bool
	Description string
}

func (t *Token) String() string {
	result := TokenPrefix + t.Name

	if t.Optional {
		result += "?"
	}

	if t.Description != "" {
		result += fmt.Sprintf(" (%s)", t.Description)
	}
	return result
}

type Flag struct {
	Data  *Token
	Value string
}

func ParseString(content string, tokens ...*Token) []*Flag {
	flags := make([]*Flag, 0)

	for _, token := range tokens {
		if strings.Contains(content, TokenPrefix+token.Name) {
			flag := &Flag{Data: token}

			arg := TokenPrefix + token.Name
			if strings.Contains(content, arg) {
				var from, to int

				from = strings.Index(content, arg) + len(arg)
				if strings.Contains(content[from:], "--") {
					to = from + strings.Index(content[from:], "--") - 1
				} else {
					to = from + len(content[from:])
				}

				flag.Value = strings.Trim(content[from:to], " ")
			}

			flags = append(flags, flag)
		}
	}

	return flags
}

func Values(flags []*Flag) []string {
	values := make([]string, 0)

	for _, flag := range flags {
		if flag.Value != "" && strings.Trim(flag.Value, " ") != "" {
			values = append(values, flag.Value)
		}
	}

	return values
}

func ResolveFlags(flags []*Flag, tokens []*Token) bool {
	if len(flags) == 0 {
		return false
	}

	for _, token := range tokens {
		if token != nil {
			if !token.Optional {
				var found bool

				for _, flag := range flags {
					if flag.Data.Name == token.Name {
						found = true
					}

					if flag.Data.Name == token.Name && strings.Trim(flag.Value, " ") == "" {
						return false
					}
				}

				if !found {
					return false
				}
			}
		}
	}

	return true
}

func GetFlag(flags []*Flag, name string) *Flag {
	for _, flag := range flags {
		if flag.Data.Name == name {
			return flag
		}
	}

	return nil
}

func HasFlag(flags []*Flag, name string) bool {
	for _, flag := range flags {
		if flag.Data.Name == name {
			return true
		}
	}

	return false
}

func GetUsages(tokens ...*Token) string {
	var result string
	for _, token := range tokens {
		result += fmt.Sprintf("%s\n", token.String())
	}

	return result
}
