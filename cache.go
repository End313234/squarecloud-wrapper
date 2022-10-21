package squarecloud

import "reflect"

type BaseCache[T any] []T

// Finds an item inside the cache structure in O(n) time, where
// n is the length of the array
func (bc *BaseCache[T]) Find(target T) int {
	for index, item := range *bc {
		if reflect.DeepEqual(target, item) {
			return index
		}
	}

	return -1
}

// Checks if an item exists inside the cache structure in O(n) time,
// where n is the length of the array
func (bc *BaseCache[T]) Contains(target T) bool {
	for _, item := range *bc {
		if reflect.DeepEqual(target, item) {
			return true
		}
	}

	return false
}

func (bc *BaseCache[T]) addsToCacheIfTargetDoesNotExist(target T) {
	if !bc.Contains(target) {
		bc.Add(target)
	}
}

// Adds an item to the cache structure
func (bc *BaseCache[T]) Add(item T) {
	*bc = append(*bc, item)
}

type UsersCache = BaseCache[User]

// Represents the client caching system
type Cache struct {
	Users UsersCache
}
