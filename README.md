# :snowflake: Echo v6

This is an ipv6 capable server. Really, it doesn't have differences with a ipv4 server. But, this service also listen in a ipv6 interface, and was made to be run in a ipv6 network.


## :bookmark_tabs: Endpoints

- GET /v1/get

Example:
```
» curl -6 localhost:3000/v1/get | jq
{
  "user-agent": "curl/7.77.0",
  "address": "[::1]:60636",
  "headers": {
    "Accept": [
      "*/*"
    ],
    "User-Agent": [
      "curl/7.77.0"
    ]
  }
}
```

## :suspect: Why? 
Usually I use [HTTPBin](https://httpbin.org) for testing and perform some basic checks in the `get` endpoint. Recently I start testing some IPV6 stuff and I realize that httpbin.org doesn't have support for it:

```
» dig -6 @8.8.8.8 httpbin.org AAAA

; <<>> DiG 9.16.15 <<>> -6 @8.8.8.8 httpbin.org AAAA
; (1 server found)
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 44159
;; flags: qr rd ra; QUERY: 1, ANSWER: 0, AUTHORITY: 1, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 512
;; QUESTION SECTION:
;httpbin.org.			IN	AAAA
```

So I made this very simple server which is capable to listen in ipv6 and ipv4 interfaces.

Feel free to reach me out.
