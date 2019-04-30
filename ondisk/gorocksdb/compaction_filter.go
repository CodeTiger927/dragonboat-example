package gorocksdb

// #include "rocksdb/c.h"
import "C"

// A CompactionFilter can be used to filter keys during compaction time.
type CompactionFilter interface {
	// If the Filter function returns false, it indicates
	// that the kv should be preserved, while a return value of true
	// indicates that this key-value should be removed from the
	// output of the compaction. The application can inspect
	// the existing value of the key and make decision based on it.
	//
	// When the value is to be preserved, the application has the option
	// to modify the existing value and pass it back through a new value.
	// To retain the previous value, simply return nil
	//
	// If multithreaded compaction is being used *and* a single CompactionFilter
	// instance was supplied via SetCompactionFilter, this the Filter function may be
	// called from different threads concurrently. The application must ensure
	// that the call is thread-safe.
	Filter(level int, key, val []byte) (remove bool, newVal []byte)

	// The name of the compaction filter, for logging
	Name() string
}

// NewNativeCompactionFilter creates a CompactionFilter object.
func NewNativeCompactionFilter(c *C.rocksdb_comparator_t) Comparator {
	return nativeComparator{c}
}

type nativeCompactionFilter struct {
	c *C.rocksdb_compactionfilter_t
}

func (c nativeCompactionFilter) Filter(level int, key, val []byte) (remove bool, newVal []byte) {
	return false, nil
}
func (c nativeCompactionFilter) Name() string { return "" }

// Hold references to compaction filters.
var compactionFilters []CompactionFilter

func registerCompactionFilter(filter CompactionFilter) int {
	compactionFilters = append(compactionFilters, filter)
	return len(compactionFilters) - 1
}
