package sized

import (
	"io"

	"github.com/cosmos/cosmos-sdk/store/types"
)

var _ types.KVStore = &Store{}

type Store struct {
	parent     types.KVStore
	sizeChange int64
}

func NewStore(parent types.KVStore) *Store {
	return &Store{parent: parent}
}
func (store *Store) GetWorkingHash() ([]byte, error) {
	return store.parent.GetWorkingHash()
}

func (s *Store) GetSizeChanged() int64 {
	return s.sizeChange
}

func (s *Store) Get(key []byte) []byte {
	value := s.parent.Get(key)
	return value
}

func (s *Store) Set(key []byte, value []byte) {
	oldValue := s.Get(key)
	if oldValue != nil {
		// reduce size due to overwrite
		s.sizeChange -= int64(len(key))
		s.sizeChange -= int64(len(oldValue))
	}
	s.parent.Set(key, value)
	if key != nil {
		s.sizeChange += int64(len(key))
	}
	if value != nil {
		s.sizeChange += int64(len(value))
	}
}

func (s *Store) Delete(key []byte) {
	// has to perform a read here to know the size change
	value := s.Get(key)
	s.parent.Delete(key)
	if value != nil {
		// only reduce size if the key used to have a value
		s.sizeChange -= int64(len(key))
		s.sizeChange -= int64(len(value))
	}
}

func (s *Store) Has(key []byte) bool {
	return s.parent.Has(key)
}

func (s *Store) Iterator(start, end []byte) types.Iterator {
	return s.parent.Iterator(start, end)
}

func (s *Store) ReverseIterator(start, end []byte) types.Iterator {
	return s.parent.ReverseIterator(start, end)
}

// GetStoreType implements the KVStore interface. It returns the underlying
// KVStore type.
func (s *Store) GetStoreType() types.StoreType {
	return s.parent.GetStoreType()
}

// CacheWrap implements the KVStore interface. It panics as a Store
// cannot be cache wrapped.
func (s *Store) CacheWrap(_ types.StoreKey) types.CacheWrap {
	panic("cannot CacheWrap a ListenKVStore")
}

// CacheWrapWithTrace implements the KVStore interface. It panics as a
// Store cannot be cache wrapped.
func (s *Store) CacheWrapWithTrace(_ types.StoreKey, _ io.Writer, _ types.TraceContext) types.CacheWrap {
	panic("cannot CacheWrapWithTrace a ListenKVStore")
}

// CacheWrapWithListeners implements the KVStore interface. It panics as a
// Store cannot be cache wrapped.
func (s *Store) CacheWrapWithListeners(_ types.StoreKey, _ []types.WriteListener) types.CacheWrap {
	panic("cannot CacheWrapWithListeners a ListenKVStore")
}
