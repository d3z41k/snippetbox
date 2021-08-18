# Snippetbox

An snippetbox in Go.

## Tips

Help

```
$ go run ./cmd/web -help
```
Environment Variables

```
$ export SNIPPETBOX_ADDR=":9999"
$ go run ./cmd/web -addr=$SNIPPETBOX_ADDR
```

Decoupled Logging

```
$ go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log
```

Disabling Directory Listings

```
$ find ./ui/static -type d -exec touch {}/index.html \;
```