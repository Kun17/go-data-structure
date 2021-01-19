package hashmap

import (
	"fmt"
	"math"

	"github.com/Kun17/go-data-structure/linkedlist"
)

// HashTable is the hash table implementation in go
type HashTable struct {
	Length int
	table  []*linkedlist.LinkedList
}

// HashData is the data we store in the bucket
type HashData struct {
	key   string
	value interface{}
}

// NewHashTable returns a new empty hash table
func NewHashTable(length int) *HashTable {
	newTable := make([]*linkedlist.LinkedList, length)
	for i := range newTable {
		newTable[i] = linkedlist.NewLinkedList()
	}
	return &HashTable{
		Length: length,
		table:  newTable,
	}
}

func calHash(val string) int {
	h := 0
	for pos, char := range val {
		h += int(char) * int(math.Pow(31, float64(len(val)-pos+1)))
	}
	return h
}

func index(hashNum int, length int) int {
	return hashNum % length
}

// Add adds a new key-value pair into the map
func (h *HashTable) Add(key string, val interface{}) error {
	i := index(calHash(key), h.Length)
	for curNode := h.table[i].Head; curNode != nil; curNode = curNode.Next {
		res, ok := curNode.Data.(HashData)
		if ok {
			if res.value == val {
				return nil
			}
		} else {
			return fmt.Errorf("Type assertion failed in hash add")
		}
	}
	h.table[i].Append(HashData{
		key:   key,
		value: val,
	})
	return nil
}

// Get returns the value by the key from the hash table
func (h *HashTable) Get(key string) (interface{}, error) {
	i := index(calHash(key), h.Length)
	for curNode := h.table[i].Head; curNode != nil; curNode = curNode.Next {
		res, ok := curNode.Data.(HashData)
		if ok {
			if calHash(res.key) == calHash(key) {
				return res.value, nil
			}
		} else {
			return 0, fmt.Errorf("Type assertion failed in hash add")
		}
	}
	return 0, fmt.Errorf("No data found")
}

// Delete deletes the key-value pair by key
func (h *HashTable) Delete(key string) error {
	i := index(calHash(key), h.Length)
	pos := 0
	for curNode := h.table[i].Head; curNode != nil; curNode = curNode.Next {
		res, ok := curNode.Data.(HashData)
		if ok {
			if calHash(res.key) == calHash(key) {
				return h.table[i].Delete(pos)
			}
		} else {
			return fmt.Errorf("Type assertion failed in hash add")
		}
		pos++
	}
	return fmt.Errorf("No such key %v found", key)
}
