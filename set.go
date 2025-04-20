/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-20 19:01:41
 * @LastEditTime: 2025-04-20 21:48:19
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /goset/set.go
 */

package goset

type Void struct{}
type Set[Key comparable] map[Key]Void

// Check existence of a key.
// Like count in CPP.
func (x Set[Key]) Has(key Key) bool {
	_, exists := x[key]
	return exists
}

// Insert key to set.
// Like insert in CPP.
func (x Set[Key]) Insert(key Key) {
	x[key] = Void{}
}

// Remove key from set.
// Like erase in CPP.
func (x Set[Key]) Erase(key Key) {
	delete(x, key)
}

// Get elem count in set.
// Like size in CPP.
func (x Set[Key]) Size() int {
	return len(x)
}
