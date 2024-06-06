package storage

type StorageI interface {
	Public() PublicI
	Party() PartyI
}

type PublicI interface {
}

type PartyI interface {
}
