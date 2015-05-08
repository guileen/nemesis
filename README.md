# Nemesis

Reverse proxy caching

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
