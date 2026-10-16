package conf

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

const (
	VERSION = 1
	PATH    = "/.taskig/"
	CONF    = "conf.toml"
)

type Config struct {
	Version int      `toml:"version"`
	Todos   []string `toml:"todos"`
}

func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	err = os.Mkdir(home+PATH, 0755)
	if err != nil {
		return err
	}

	conf := Config{
		Version: 1,
	}

	b, err := toml.Marshal(conf)
	if err != nil {
		return err
	}

	// Create an empty config
	err = os.WriteFile(home+PATH+CONF, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	b, err := os.ReadFile(home + PATH + CONF)
	if err != nil {
		return nil, err
	}

	var conf Config
	toml.Unmarshal([]byte(b), &conf)

	return &conf, nil
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return home + PATH + CONF, nil
}

func ListTodos() error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		if conf.Version > VERSION {
			return fmt.Errorf("config version doesn't match, i'm too old!")
		} else {
			return fmt.Errorf("config version doesn't match, i'm too new!")
		}
	}

	for i, todo := range conf.Todos {
		fmt.Printf("%3d: %s\n", i, todo)
	}

	return nil
}

func AddTodoTask(name string) error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		if conf.Version > VERSION {
			return fmt.Errorf("config version doesn't match, i'm too old!")
		} else {
			return fmt.Errorf("config version doesn't match, i'm too new!")
		}
	}

	// Add a new task
	conf.Todos = append(conf.Todos, name)

	// Get config path
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	// Save it
	b, err := toml.Marshal(conf)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func RmTodoTask(name string) error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		if conf.Version > VERSION {
			return fmt.Errorf("config version doesn't match, i'm too old!")
		} else {
			return fmt.Errorf("config version doesn't match, i'm too new!")
		}
	}

	// If there are no task we can already just return
	if len(conf.Todos) == 0 {
		return fmt.Errorf("%q doesn't exist", name)
	}

	new_todos := make([]string, 0, len(conf.Todos)-1)
	for _, el := range conf.Todos {
		if el == name {
			continue
		} else {
			new_todos = append(new_todos, el)
		}
	}

	// If no elements got removed, then error
	if len(new_todos) >= len(conf.Todos) {
		return fmt.Errorf("%q doesn't exist", name)
	}

	conf.Todos = new_todos

	// Get config path
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	// Save it
	b, err := toml.Marshal(conf)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func RmTodos() error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		if conf.Version > VERSION {
			return fmt.Errorf("config version doesn't match, i'm too old!")
		} else {
			return fmt.Errorf("config version doesn't match, i'm too new!")
		}
	}

	conf.Todos = []string{}

	// Get config path
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	// Save it
	b, err := toml.Marshal(conf)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
