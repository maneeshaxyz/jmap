package main

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{map[string]string{}}
}

type InMemoryUserStore struct {
	store map[string]string
}

func (i *InMemoryUserStore) GetUserString(name string) string {
	return i.store[name]
}

func (i *InMemoryUserStore) ChangeUserValues(name string) {
	i.store[name] = "changed value"
}
