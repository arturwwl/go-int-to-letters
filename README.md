# GO INT TO LETTERS

Why should You need it? For instance to get Excel column name by index.


###  example

```go
package main

import (
	"fmt"
	gointtoletters "github.com/arturwwl/gointtoletters"
)

func main() {
	fmt.Println(gointtoletters.IntToLetters(1))// print A
	fmt.Println(gointtoletters.IntToLetters(26))// print Z
	fmt.Println(gointtoletters.IntToLetters(27))// print AA
	fmt.Println(gointtoletters.IntToLetters(1996))// print BXT
	fmt.Println(gointtoletters.IntToLetters(787428))// print ARTUR

	return
}
```
