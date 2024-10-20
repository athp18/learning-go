package Hashmap

type Entry struct {
	key   string
	value interface{}
	next  *Entry
}

type HashMap struct {
	buckets []*Entry
	size    int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([]*Entry, size),
		size:    size,
	}
}

func (hm *HashMap) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = 31*hash + int(key[i])
	}
	return hash % hm.size
}

func (hm *HashMap) Put(key string, value interface{}) {
	index := hm.hash(key)
	entry := hm.buckets[index]

	if entry == nil {
		// No collision, insert new entry
		hm.buckets[index] = &Entry{key, value, nil}
	} else {
		// Handle collision with separate chaining (linked list)
		for entry != nil {
			if entry.key == key {
				entry.value = value // Update existing value
				return
			}
			if entry.next == nil {
				entry.next = &Entry{key, value, nil}
				return
			}
			entry = entry.next
		}
	}
}

func (hm *HashMap) Get(key string) (interface{}, bool) {
	index := hm.hash(key)
	entry := hm.buckets[index]

	for entry != nil {
		if entry.key == key {
			return entry.value, true
		}
		entry = entry.next
	}

	return nil, false
}

func (hm *HashMap) Remove(key string) bool {
	index := hm.hash(key)
	entry := hm.buckets[index]

	if entry == nil {
		return false
	}

	if entry.key == key {
		hm.buckets[index] = entry.next
		return true
	}

	prev := entry
	for entry != nil {
		if entry.key == key {
			prev.next = entry.next
			return true
		}
		prev = entry
		entry = entry.next
	}

	return false
}
