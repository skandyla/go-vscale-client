# example vscale api client

### motivation
- get insite how to write sdk for real projects
- use interfaces (for SOLID, mocking)
- use subServices structure I.e.  client.Droplet.List(), where List method can be defined for multiple services
- write something useful for yourself (vscale is cheapest option) for testing misc automations
- project structure is inspired by Digital Ocean go client


### testing
```
$ go test -v .
=== RUN   TestRegions_List
--- PASS: TestRegions_List (0.00s)
PASS
ok      github.com/skandyla/go-vscale-client    (cached)

$ go test -v -tags="integration" ./test/integration/...
=== RUN   TestRegionsService_List
[Fri Nov  5 08:16:36 2021] GET https://api.vscale.io/v1/locations
--- PASS: TestRegionsService_List (1.36s)
PASS
ok      github.com/skandyla/go-vscale-client/test/integration   (cached)
```