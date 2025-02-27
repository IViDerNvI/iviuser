package store

type Store interface {
	Users() UserStore
	Close()
}

var cli *Store

func Factory() Store {
	return *cli
}

func SetFactory(store Store) {
	cli = &store
}
