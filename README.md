# ltsvlog

**Warning** Under developing. API will be changed in further commits.

## Examples

### with struct

```
package main

import . "github.com/harukasan/ltsvlog"

type L struct{
  Time time.Time
  Status int
  URL string
}

func main() {
  Log(L{Time: time.Now, Status: 200, URL: "https://example.com/"})
  // time:1970-01-01T09:00:00+09:00\tstatus:200\turl:https://example.com/
}
```

### with fields

```
package main

import . "github.com/harukasan/ltsvlog"

func main() {
  Logf(F("time", time.Now), F("status", 200), F("url", "https://example.com/"))
  // time:1970-01-01T09:00:00+09:00\tstatus:200\turl:https://example.com/
}
```

### Copyright

Copyright 2017 Shunsuke Michii. All rights reserverd.
See [LICENSE](./LICENSE.md) file.
