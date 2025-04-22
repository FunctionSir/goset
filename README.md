<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-20 18:58:16
 * @LastEditTime: 2025-04-22 16:08:43
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /goset/README.md
-->
# goset

CPP-like set in Golang.

Warning: this is NOT a ordered set!

Well tested.

## How to use

``` go
import "github.com/FunctionSir/goset"
// ... //
someSet := make(Set[int])
someSet.Insert(1)
if someSet.Has(1) {
    // ... //
}
someSet.Erase(1)
if someSet.Size() == 0 {
    // ... //
}
```
