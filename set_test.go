/*
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-20 19:09:56
 * @LastEditTime: 2025-04-20 21:54:52
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /goset/set_test.go
 */

package goset

import (
	"testing"
)

func TestSet(t *testing.T) {
	sA := make(Set[int])
	sA.Insert(1)
	sA.Insert(2)
	if !sA.Has(1) || !sA.Has(2) {
		t.Errorf("Should have 1 and 2 in set!")
	}
	if sA.Has(3) {
		t.Errorf("Should have no 3 in set!")
	}
	if sA.Size() != 2 {
		t.Errorf("Size should be 2!")
	}
	sA.Erase(1)
	sA.Erase(2)
	if sA.Has(1) || sA.Has(2) {
		t.Errorf("Should have no 1 or 2 in set!")
	}
	if sA.Size() != 0 {
		t.Errorf("Size should be 0!")
	}
	sB := make(Set[string])
	sB.Insert("A")
	sB.Insert("a")
	if !sB.Has("A") || !sB.Has("a") {
		t.Errorf("Should have \"A\" and \"a\" in set!")
	}
	if sB.Has("B") {
		t.Errorf("Should have no \"B\" in set!")
	}
}
