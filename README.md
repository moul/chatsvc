# chatsvc
:gift: chat micro-service using nats and go-kit (gRPC + http)

A chat micro-service, developed using [nats](http://nats.io), [Protobuf](https://github.com/google/protobuf) and [Go-Kit](https://github.com/go-kit/kit), with more that 75% boilerplate code generated automatically using [protoc-gen-gotemplate](https://github.com/moul/protoc-gen-gotemplate).

## Usage

Server

```console
$ chatsvc
ts=2016-12-27T12:49:44Z caller=main.go:62 transport=gRPC addr=:9000
Received a message from "pid32686": "hello world"
Received a message from "pid32700": "hello pid32686"{"offset":15}
```

---

Client 1

```console
$ ./chatsvc-client
hello world
pid32700> hello pid32686
```

---

Client 2

```console
$ ./chatsvc-client
pid32686> hello world
hello pid32686
```

---

Client 3

```console
$ ./chatsvc-client
pid32686> hello world
pid32686> hello world
```
