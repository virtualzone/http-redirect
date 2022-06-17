# HTTP Redirect
**A tiny HTTP server written in Go which does one thing: Redirect incoming requests to another URL.**

## Quick start
To run with Docker and redirect incoming requests to `https://virtualzone.de`, use:

```bash
docker run --rm -d \
    -p 8080:8080 \
    -e TARGET="https://virtualzone.de" \
    virtualzone/http-redirect
```

## Environment variables
Use environment variables to configure HTTP Redirect:

Env | Default | Description
--- | --- | ---
LISTEN_ADDR | 0.0.0.0:8080 | TCP Listening address and port
TARGET | &lt;empty&gt; | http://localhost/
APPEND_PATH | 1 | Append path from requested URL to redirect URL
STATUS_CODE | 301 | HTTP Status Code
