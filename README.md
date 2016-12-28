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

As a webserver

```
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/0x4445565a/nice-things"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	token := "Meranda"
	if val, ok := q["token"]; ok {
		token = val[0]
	}
	s := niceThings.GeneratePhrase()
	fmt.Fprintf(w, "%s", strings.Replace(s, "{0}", token, -1))

}

/**
 * Create server at localhost:8080 use ?token=value to
 * set token.
 */
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```
