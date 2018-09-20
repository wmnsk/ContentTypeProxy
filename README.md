# ContentTypeProxy

Dead simple HTTP Proxy that inserts Content-Type if the Content-Type field is missing or empty.
This is just a PoC to show how easy it is to edit HTTP Header with Proxy.

Please note that this is NOT for production, even it might also be the help for someone in trouble with HTTP Header :sweat_smile:

https://seclists.org/fulldisclosure/2018/Sep/16

## Special Thanks

Large part of the code that make it work as HTTP proxy is retrieved from [HTTP(S) Proxy in Golang in less than 100 lines of code](https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c) by [@mlowicki](https://medium.com/@mlowicki).

## Installing

Download `main.go` and build it, and you can get the executable.

```shell-session
git clone git@github.com:wmnsk/ContentTypeProxy.git
cd ContentTypeProxy
go build
```

## Usage

### Run

```shell-session
Usage of ContentTypeProxy:
  -addr string
        address to serve HTTP (default "0.0.0.0:55555")
  -type string
        Content-Type to set (default "text/plain")
```

### Use

Set your browser's proxy to the address:port you specified in the command-line args above.

## Author

Yoshiyuki Kurauchi ([GitHub](https://github.com/wmnsk/) / [Twitter](https://twitter.com/wmnskdmms))

## License

[MIT](https://github.com/wmnsk/ContentTypeProxy/blob/master/LICENSE)
