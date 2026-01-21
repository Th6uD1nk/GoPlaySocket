# RTGS Client

## Run the client

```bash
go run rtgs-client
````

CGO_CPPFLAGS="-v" go run -x ./...
gomobile build -target=android -androidapi=21 -tags=mobile .

## Build the client

```bash
CGO_CPPFLAGS="-v" go build -x ./...
```

