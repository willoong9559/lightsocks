# lightsocks

An experimental implementation of shadowsocks

# Features

**It's simple but stable, with a simplified forwarding architecture to implementing basic functions. Does not amplify the transmission traffic, which is less and faster and takes up fewer resources (memory usage of about 1M and CPU usage of about 0.1%).**

# Build

## Requirements

+ git

+ go 1.19+

## Build Steps

```bash

# clone and enter the repo

make

# or `make server/client' to build erver/client separately

```

Execute the binary file

# Usage

Generate a psk for encryption
```
./server -g
```

An ini file is needed:

```ini
# section "client"
[client]
listen = 0.0.0.0:1234
server = 1.2.3.4:5678
psk = psk

# section "server"
[server]
listen = 0.0.0.0:5678
psk = psk
```

```bash
./server -c ./server.conf
./client -c ./client.conf
```

# Thanks

+ [gwuhaolin/lightsocks](https://github.com/gwuhaolin/lightsocks)

