# go-radius

[![GoDoc](https://godoc.org/github.com/blind-oracle/go-radius?status.svg)](https://godoc.org/github.com/blind-oracle/go-radius)
[![cover.run](https://cover.run/go/github.com/blind-oracle/go-radius.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fblind-oracle%2Fgo-radius)
[![Go Report Card](https://goreportcard.com/badge/github.com/blind-oracle/go-radius)](https://goreportcard.com/report/github.com/blind-oracle/go-radius)

It's quite heavily rewritten fork of another Go [RADIUS library](https://github.com/layeh/radius)

Significant changes are:
* Common
  * Encoding/Decoding of attribute 26 (Vendor-Specific)
  * RFC2866 & RFC2869 (Accounting)
+ 
  * Support Vendor Specific Dictionaries
  * VSA can integrates to any dictionary (support rfc and VSA dictionaries together)
  * Transform all attribute types from string values (include VSA)
  * Support Tagged attributes
  * Encodes and Decodes VSA and VSA Tagged attribute values
  
* Server
  * Request throttling (maximum requests per second) support
  * Supports limiting the number of requests in processing queue
  * Multiple RADIUS Secrets based on packet's source IP with a fallback default
  * Request/Response packet replication (useful for logging, IDS etc)
  * Configurable UDP buffer size

* Client
  * Lots of vendor-specific (Cisco, Juniper, Mikrotik) functions and constants
  * Support for generating CoA/Disconnect-Message packets

## Installation
    go get -u github.com/blind-oracle/go-radius

## Server example
```go
import (
    "github.com/blind-oracle/go-radius"
    "log"
)

func main() {
    handler := func (w radius.ResponseWriter, p *radius.Packet) {
        w.AccessAccept()
    }

    server := radius.Server{
        Addr:           "0.0.0.0:1812",
        Handler:        radius.HandlerFunc(handler),
        Secret:         []byte(o.RADIUSSecret),
        Dictionary:     radius.Builtin,
    }

    if err := server.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
```

## Client example
```go
import (
    "github.com/blind-oracle/go-radius"
    "log"
)

func main() {
    client := radius.Client{}
    packet := radius.New(radius.CodeAccessRequest, []byte("VerySecret"))
    packet.Add("Calling-Station-Id", "NAS-Fake")

    reply, err := client.Exchange(packet, "1.2.3.4:1812")
    if err != nil {
        log.Fatalf(err)
    }

    switch reply.Code {
        case radius.CodeAccessAccept:
        log.Println("Accept")
        case radius.CodeAccessReject:
        log.Println("Reject")
    }
}
```

## Authors
* Tim Cooper (<tim.cooper@layeh.com>)
* Igor Novgorodov (<igor@novg.net>)
* Andrey Melnikov (<echo@orn.ru>)