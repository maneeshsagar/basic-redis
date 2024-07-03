package persistence

type Persistence interface {
	PutPersistence(key string, value interface{}) (err error)
	PutTimeoutPersistence(key string, value interface{}, timeout int) (err error)
	GetPersistence(key string) (value string, err error)
}

type RedisPersistence struct {
}

func NewRedisPersistence() Persistence {
	return &RedisPersistence{}
}
