# go-instagram

[![Build Status](https://travis-ci.org/hieven/go-instagram.svg?branch=master)](https://travis-ci.org/hieven/go-instagram)[![codecov](https://codecov.io/gh/hiEven/go-instagram/branch/master/graph/badge.svg)](https://codecov.io/gh/hiEven/go-instagram)

This project is for study and personal use only. We hold no responsibilty of any use that violate Instagram's terms and conditions.

## Installation

```sh
$ go get github.com/hieven/go-instagram
```

## Documentation

- [![GoDoc](https://godoc.org/github.com/hieven/go-instagram?status.svg)](https://godoc.org/github.com/hieven/go-instagram) Instagram
- [![GoDoc](https://godoc.org/github.com/hieven/go-instagram/models?status.svg)](https://godoc.org/github.com/hieven/go-instagram/models) Instagram Models

## Features

You can use this repo to:
- like/unlike media
- get timeline feed
- get ranked media / recent media of a location
- get inbox messages
- broadcast to any inbox thread
- approve pending inbox thread

The project is still in its early stage. Any pull request to extend its functionalities is most welcome.

## Example

```go
cnf := &config.Config{
  Username: "USERNAME",
  Password: "PASSWORD",
}

ctx := context.Background()

ig, _ := instagram.New(cnf)
ig.Login(ctx)

Aresp, _ := ig.Timeline().Feed(ctx, instagram.TimelineFeedReques{})
fmt.Println(Aresp.Items[0].MediaOrAd)
```

Find more complex examples on [examples](examples)

## License

MIT

## Similar Projects

- [instagram-private-api](https://github.com/huttarichard/instagram-private-api) (Node.js)
- [Instagram-API](https://github.com/mgp25/Instagram-API) (PHP)
