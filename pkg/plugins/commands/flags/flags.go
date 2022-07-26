package flags

import (
	"fmt"
	"strings"
)

const ShortTokenPrefix string = "-"
const TokenPrefix string = "--"

// Token represents a flag token. Like --subcommand or -sc.
type Token struct {
	Name        string
	Short       string
	Optional    bool
	Default     bool
	Description string
}

func (t *Token) String() string {
	result := TokenPrefix + t.Name

	if t.Optional {
		result += "?"
	}

	if t.Short != "" {
		result += fmt.Sprintf(" %s%s", ShortTokenPrefix, t.Short)

		if t.Optional {
			result += "?"
		}
	}

	if t.Description != "" {
		result += fmt.Sprintf(" (%s)", t.Description)
	}
	return result
}

// Flag represents a flag command line argument.
type Flag struct {
	// Data is the token that represents the flag.
	Data *Token
	// Value of the flag taken from the command line argument.
	Value string
}

// ParseString parses a string and returns a list of flags. It can be empty.
func ParseString(content string, tokens ...*Token) []*Flag {
	flags := make([]*Flag, 0)

	for _, token := range tokens {
		shortFlag := ShortTokenPrefix + token.Short + " "
		longFlag := TokenPrefix + token.Name + " "

		// If short flag is present, replace it with long flag.
		if token.Short != "" {
			if strings.Contains(content, shortFlag) {
				content = strings.Replace(content, shortFlag, longFlag, 1)
				content = strings.Trim(content, " ")
			}
		}

		// Checking the flag
		if strings.Contains(content, longFlag) {
			flag := &Flag{Data: token}
			arg := TokenPrefix + token.Name

			// If the flag exists add the value.
			if strings.Contains(content, arg) {
				var from, to int

				// Catching the from:to range value of the flag.
				from = strings.Index(content, arg) + len(arg)
				if strings.Contains(content[from:], "-") {
					to = from + strings.Index(content[from:], "-") - 1
				} else {
					to = from + len(content[from:])
				}

				flag.Value = strings.Trim(content[from:to], " ")
			}

			flags = append(flags, flag)
		}

		// If the flag can be catched without -- or -.
		if token.Default {
			if !strings.Contains(content, TokenPrefix+token.Name+" ") || !strings.Contains(content, ShortTokenPrefix+token.Short+"") {
				flag := &Flag{Data: token}
				flag.Value = content
				flags = append(flags, flag)
			}
		}
	}

	return flags
}

// Values returns the values of the flags.
func Values(flags []*Flag) []string {
	values := make([]string, 0)
	for _, flag := range flags {
		values = append(values, flag.Value)
	}

	return values
}

// ResolveFlags checks if the flags are all valid.
func ResolveFlags(flags []*Flag, tokens ...*Token) bool {
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

// GetFlag returns the flag with the given name or nil.
func GetFlag(flags []*Flag, name string) *Flag {
	for _, flag := range flags {
		if flag.Data.Name == name {
			return flag
		}
	}

	return nil
}

// HasFlag returns true if the flag is present.
func HasFlag(flags []*Flag, name string) bool {
	for _, flag := range flags {
		if flag.Data.Name == name {
			return true
		}
	}

	return false
}

// GetUsages returns the usage string for the given flags.
// It's the same as use token.String() for each token.
func GetUsages(tokens ...*Token) string {
	var result string
	for _, token := range tokens {
		result += fmt.Sprintf("%s\n", token.String())
	}

	return result
}
