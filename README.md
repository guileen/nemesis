# Nemesis

Reverse proxy caching

# install
```
$ go get github.com/guileen/nemesis
$ nemesis
```

## Config Example

```yaml
api.example.com
   /v1:
      get *:
      proxy-pass: http://127.0.0.1:3000
```

## Features

* Domain config
* URL rewrite
* Static serve
* Frequent limit
* Lua support
* Reverse proxy
    * IP hashing
    * Upstream hashing
    * Upstream compress
    * Upstream caching
* Caching
    * Memory caching
    * Local LevelDB caching
    * Remote Redis caching
* Cache purge by tag
    * If some resource has been forbiden, the busyness server can purge the cached content on the server.

## How it works

```
while New or Rewrited:
    Route
    Limiting
    Caching
    Rewrite
Handle
```

Handlers

* Upstream Handler
* Static Handler
* Script Handler

## Config example

```
upstream:
    backend:
        - http://127.0.0.1:3000/

www.thel.co:
    default: true
    host:
        - www.thel.co
    port: 80
    routes:
        - get /static/*:
            url: /$1
            static: /path/to/static/root
        - * /apiv1/*:
            url: /$1
            proxy: backend
        - *:
            status: 404
            text: Not Found
api.thel.co:
    proxy:
        - http://127.0.0.1:4000/
        - http://127.0.0.1:4001/
        - http://127.0.0.1:4002/

static.thel.co:
    static: /path/to/static/root
```

## Server

```
domain.name:
    [server options]
    [modules]
```
* default
* host
* port
* modules

## Modules

### routes

```
routes:
    - <method> <path>:
        [modules]
    ...
```

*

### url

```
url: <expression>
```

$1

### status

`status: <status>`

Set response status.

### text

`text: <text>`

Set response text.

### proxy

```
proxy:
    - <upstream_url>
```

Proxy resonpse to upstream node.

### static

```
static: <root_path>
```

Serve static files.
