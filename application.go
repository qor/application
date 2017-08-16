package application

// Application application structure
type Application struct {
	Config *Config
}

// Config application config
type Config struct {
}

// New new application
func New(config *Config) *Application {
	return &Application{
		Config: config,
	}
}
