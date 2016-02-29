sanic
=====

sanic is a clone of
[Twitter snowflake](https://github.com/twitter/snowflake/tree/snowflake-2010)
(the 2010 version), written in [Golang](https://golang.org/).
More specifically, the [IdWorker section of snowflake]
(https://github.com/twitter/snowflake/blob/snowflake-2010/src/main/scala/com/twitter/service/snowflake/IdWorker.scala).

### Usage

To use sanic, either make a new worker, or select a premade one:

```go
package main

import (
	"fmt"

	"github.com/ifo/sanic"
)

func main() {
	worker := sanic.SevenLengthWorker
	// equivalent to:
	// worker := sanic.NewWorker(0, 1451606400, 0, 10, 31, time.Second)

	id := worker.NextID()
	idString := worker.IDString(id)
	fmt.Println(id)       // e.g. 5292179457
	fmt.Println(idString) // e.g. "AUBwOwE"
}
```

Check out [the examples](https://github.com/ifo/sanic/tree/master/examples) for
more.

## License

sanic is ISC licensed.
Check out the [LICENSE](https://github.com/ifo/sanic/blob/master/LICENSE) file.
