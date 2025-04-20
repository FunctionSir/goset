<!--
 * @Author: FunctionSir
 * @License: AGPLv3
 * @Date: 2025-04-20 18:58:16
 * @LastEditTime: 2025-04-20 21:53:36
 * @LastEditors: FunctionSir
 * @Description: -
 * @FilePath: /goset/README.md
-->
# goset

CPP-like set in Golang.

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
