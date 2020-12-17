# FalconDB
```

  ______    _                 _____  ____  
 |  ____|  | |               |  __ \|  _ \ 
 | |__ __ _| | ___ ___  _ __ | |  | | |_) |
 |  __/ _` | |/ __/ _ \| '_ \| |  | |  _ < 
 | | | (_| | | (_| (_) | | | | |__| | |_) |
 |_|  \__,_|_|\___\___/|_| |_|_____/|____/ 
                               
```

FalconDB is an in-memory key-value databse written in Golang.
It supports [redis](https://redis.io/) protocol for these listed commands:
- GET
- SET
- DEL

Example:
Please provide your desired port as an input argument.

``` bash
./falconDB 8585

SET foo bar
OK!
GET foo
bar
DEL foo
OK!
GET foo
<nil>
```
