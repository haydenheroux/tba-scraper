package scout

type Scout struct {
	URL string
}

func New(dbAddress string) Scout {
	return Scout{
		URL: dbAddress,
	}
}
