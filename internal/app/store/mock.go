package store

import "log"

type Store interface {
	Open() error
	Close()
	GetItem(key string) (string, bool)
	PutItem(key, item string)
}

type MemoryStore struct {
	data map[string]string
}

func NewStore() *MemoryStore {
	return &MemoryStore{
		make(map[string]string),
	}
}

func (store *MemoryStore) Open() error {
	log.Println("connected to mock database")
	return nil
}

func (store *MemoryStore) Close() {
	log.Println("mock database closed")
}

func (store *MemoryStore) GetItem(shortUrl string) (string, bool) {
	sourceUrl, ok := store.data[shortUrl]
	if !ok {
		return "", false
	}
	return sourceUrl, true
}

func (store *MemoryStore) PutItem(shortUrl, sourceUrl string) {
	store.data[shortUrl] = sourceUrl
}
