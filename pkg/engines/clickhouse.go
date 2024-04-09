package engines

type ClickhouseEngine struct {
	host     string
	port     int
	database string
	username string
	password string
}

func (c *ClickhouseEngine) Publish(data map[string]any) error {
	return nil
}

func NewClickhouseEngine(host string, port int, database, username, password string) ClickhouseEngine {
	return ClickhouseEngine{
		host:     host,
		port:     port,
		database: database,
		username: username,
		password: password,
	}
}
