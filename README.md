# go-ts3-http

### usage
see `example/example.go`

### available endpoints (from `docs`)
Endpoints as described in [here](https://community.teamspeak.com/t/webquery-discussion-help-3-12-0-onwards/7184).

- ❌ `not started`
- ✔ `implemented`
- ⚠ `work in progress`

| endpoint doc file | status |
|:---|:---:|
| apikey | ✔ |
| ban | ✔ |
| channel | ✔ |
| channelclientperm | ⚠ |
| channelgroup | ⚠ |
| channelperm | ✔ |
| client | ✔ |
| complain | ✔ |
| custom | ✔ |
| general | ✔ |
| instance | ✔ |
| log | ✔ |
| message | ✔ |
| permission | ✔ |
| privilegekey | ✔ |
| querylogin | ✔ |
| server | ✔ |
| servergroup | ✔ |
| servertemppassword | ✔ |
| token | ✔ |

### TODO
- [ ] implement all endpoints
- [x] implement an event system based on raw (until events for the http api gets announced)