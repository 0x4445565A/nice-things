### Usage ###
```
package main

import (
	"fmt"
	"strings"

	"github.com/0x4445565a/nice-things"
)

func main() {
	s := niceThings.GeneratePhrase()
	fmt.Println(strings.Replace(s, "{0}", "Something", -1))
}
```
