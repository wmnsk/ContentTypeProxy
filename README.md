# ContentTypeProxy

Dead simple HTTP Proxy that inserts Content-Type if the Content-Type field is missing or empty in the header of POST.
This is just a PoC to show how easy it is to edit HTTP Header with Proxy.

Please note that this is NOT for production, even though it might also be the help for someone in trouble with HTTP Header :sweat_smile:

https://seclists.org/fulldisclosure/2018/Sep/16

## Special Thanks

Large part of the code that make it work as HTTP proxy is retrieved from [HTTP(S) Proxy in Golang in less than 100 lines of code](https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c) by [@mlowicki](https://medium.com/@mlowicki).

## Installing

### Use Executables

Executables can be found in [releases](https://github.com/wmnsk/ContentTypeProxy/releases) page.

|         |                                                                                                                 |
| ------- | --------------------------------------------------------------------------------------------------------------- |
| Linux   | [ContentTypeProxy](https://github.com/wmnsk/ContentTypeProxy/releases/download/v0.1.0/ContentTypeProxy)         |
| Windows | [ContentTypeProxy.exe](https://github.com/wmnsk/ContentTypeProxy/releases/download/v0.1.0/ContentTypeProxy.exe) |

### Build yourself

Download `main.go` and build it, and you can get the executable.

```shell-session
git clone git@github.com:wmnsk/ContentTypeProxy.git
cd ContentTypeProxy
go build
```

## Usage

### Run

Run ContentTypeProxy on a server, and set your browser's proxy to that server.

Address:Port to listen and Content-Type to be added can be specified in command-line arguments.

```shell-session
Usage of ContentTypeProxy:
  -addr string
        address to serve HTTP (default "0.0.0.0:55555")
  -type string
        Content-Type to set (default "text/plain")
```


## Author

Yoshiyuki Kurauchi ([GitHub](https://github.com/wmnsk/) / [Twitter](https://twitter.com/wmnskdmms))

## License

[MIT](https://github.com/wmnsk/ContentTypeProxy/blob/master/LICENSE)
