# example vscale api client

### motivation
- get insite how to write sdk for real projects
- use interfaces (for SOLID, mocking)
- use subServices structure I.e.  client.Droplet.List(), where List method can be defined for multiple services
- write something useful for yourself (vscale is cheapest option) for testing misc automations
- project structure is inspired by Digital Ocean go client

# example output:
```
[Wed Nov  3 20:13:25 2021] GET https://api.vscale.io/v1/locations
Locations: spb0
Locations: msk0
[Wed Nov  3 20:13:26 2021] GET https://api.vscale.io/v1/scalets
Scalet: Stunning-Jangala 03.11.2021 15:54:04 spb0 84.38.183.125
[Wed Nov  3 20:13:26 2021] POST https://api.vscale.io/v1/scalets
Created scalet: {6211632 test queued spb0 small [{baikonur-ed25519 53551}] [] {  } map[] debian_10_64_001_master  03.11.2021 17:13:34 true true <nil> <nil> <nil> <nil>}
Scalet: test 03.11.2021 17:13:34 
```