package store

type Store interface {
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
