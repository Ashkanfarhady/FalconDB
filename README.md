# FalconDB

```

  ______    _                 _____  ____
 |  ____|  | |               |  __ \|  _ \
 | |__ __ _| | ___ ___  _ __ | |  | | |_) |
 |  __/ _` | |/ __/ _ \| '_ \| |  | |  _ <
 | | | (_| | | (_| (_) | | | | |__| | |_) |
 |_|  \__,_|_|\___\___/|_| |_|_____/|____/

```

FalconDB is an in-memory key-value database written in Golang.
It supports [redis](https://redis.io/) protocol for these listed commands:

- GET
- SET
- DEL

Example:
Please provide your desired port as an input argument.

``` bash
./falconDB 8585

redis-cli -p 8585 SET foo bar
OK
redis-cli -p 8585 GET foo
bar
redis-cli -p 8585 DEL foo
OK
```

You can also use [telnet](https://en.wikipedia.org/wiki/Telnet) to connect to FalconDB:

```bash
telnet localhost 8585


*3
$3
SET
$3
foo
$3
bar
```

- `*3` indicates there are 3 parts to this command (SET, key, and value).
- `$3` indicates the next line has 3 characters (the word "SET").
- `SET` is the command.
- `$3` again indicates the key is 3 characters long ("foo").
- `foo` is the actual key.
- `$3` indicates the value is 3 characters long ("bar").
- `bar` is the actual value.
