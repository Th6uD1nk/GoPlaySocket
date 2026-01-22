# RTGS Client

## Run the client

```bash
go run rtgs-client
````

or

```bash
CGO_CPPFLAGS="-v" go run -x ./...
```

## Build the client for pc

```bash
CGO_CPPFLAGS="-v" go build -x -o rtgs-client
```

## Build the client for mobile

```bash
gomobile build -target=android -androidapi=33 -tags=mobile .
```

## Install and run client on Android + logs
```bash
adb shell am force-stop com.th6ud1nk.rtgsclient && \
adb install -r ./rtgs_client.apk && \
adb push ./rtgs-config.txt /sdcard/Download/rtgs-config.txt && \
adb shell am start -n com.th6ud1nk.rtgsclient/org.golang.app.GoNativeActivity && \
adb logcat | grep GoLog
```
