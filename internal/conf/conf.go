package conf

import (
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/pelletier/go-toml/v2"
)

const (
	VERSION = 2
	PATH    = "/.taskig/"
	CONF    = "conf.toml"
)

type Event struct {
	Begin time.Time
	End   time.Time
	Name  string
}

type Config struct {
	Version int      `toml:"version"`
	Todos   []string `toml:"todos"`
	Events  []Event  `toml:"events"`
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
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
		}
	}

	for i, todo := range conf.Todos {
		fmt.Printf("%3d: %s\n", i, todo)
	}

	return nil
}

func TidyEvents() error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
		}
	}

	current := time.Now()

	new_events := make([]Event, 0)
	for _, ev := range conf.Events {
		if current.After(ev.End) {
			continue
		} else {
			new_events = append(new_events, ev)
		}
	}

	// Replace the old events with the cleaned up ones
	conf.Events = new_events

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

func ListEvents() error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
		}
	}

	current := time.Now()

	// Sort everything by time
	events := slices.Clone(conf.Events)
	slices.SortFunc(events, func(a, b Event) int {
		return a.Begin.Compare(b.Begin)
	})

	i := 0
	for _, ev := range events {
		if !ev.End.After(current) {
			continue // Already finished
		}

		if ev.Begin.Before(current) && ev.End.After(current) {
			i++
			fmt.Printf("%3d: %s (ends in %s) from %s to %s\n",
				i, ev.Name, ev.End.Sub(current).Round(time.Second), ev.Begin.Format("2006-Jan-02 15:04"), ev.End.Format("2006-Jan-02 15:04"))
		} else if ev.Begin.After(current) {
			i++
			fmt.Printf("%3d: %s (starts in %s) from %s to %s\n",
				i, ev.Name, ev.Begin.Sub(current).Round(time.Second), ev.Begin.Format("2006-Jan-02 15:04"), ev.End.Format("2006-Jan-02 15:04"))
		}
	}

	return nil
}

func ListAllEvents() error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
		}
	}

	events := slices.Clone(conf.Events)
	slices.SortFunc(events, func(a, b Event) int {
		return a.Begin.Compare(b.Begin)
	})

	for i, ev := range events {
		fmt.Printf("%3d: %s from %s to %s\n",
			i, ev.Name, ev.Begin.Format("2006-Jan-02 15:04"), ev.End.Format("2006-Jan-02 15:04"))
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
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
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
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
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

func RmEvents() error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
		}
	}

	conf.Events = []Event{}

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
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
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

func AddEvent(event Event) error {
	// Get the config
	conf, err := getConfig()
	if err != nil {
		return err
	}

	// Validate the version
	if conf.Version != VERSION {
		switch {
		case conf.Version < VERSION:
			return fmt.Errorf("config version %d is older than supported version %d, i'm too new!", conf.Version, VERSION)
		case conf.Version > VERSION:
			return fmt.Errorf("config version %d is newer than supported version %d, i'm too old!", conf.Version, VERSION)
		}
	}

	conf.Events = append(conf.Events, event)

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
