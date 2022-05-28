# Pequod

## Named for the boat that chased Moby-Dick

This repository is used to store protobuf definitions and common business logic between [marina](https://github.com/skuid/marina) and [tides](https://github.com/skuid/tides)

We're going to use [buf](https://buf.build)([pricing](https://buf.build/pricing/)) for the protocol buffer builds if we want to later.

Until then, makefiles.

---

## proto deps

```
brew install protobuf
go install github.com/favadi/protoc-go-inject-tag
```

---

## here's the buf stuff

### Update your crap (mac)

```
softwareupdate --all --install --force
```


### Install `buf`

```
brew install bufbuild/buf/buf
```

### Install `jq`

```
brew install jq
```

### Install `bloomrpc`

```
brew install bloomrpc
```