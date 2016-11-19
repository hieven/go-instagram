# go-instagram

[![Build Status](https://travis-ci.org/hieven/go-instagram.svg?branch=master)](https://travis-ci.org/hieven/go-instagram)[![codecov](https://codecov.io/gh/hiEven/go-instagram/branch/master/graph/badge.svg)](https://codecov.io/gh/hiEven/go-instagram)

## Installation

```sh
$ go get github.com/hieven/go-instagram
```

## Documentation

- [![GoDoc](https://godoc.org/github.com/hieven/go-instagram?status.svg)](https://godoc.org/github.com/hieven/go-instagram) Instagram
- [![GoDoc](https://godoc.org/github.com/hieven/go-instagram/models?status.svg)](https://godoc.org/github.com/hieven/go-instagram/models) Instagram Models

## Features

This is still in an early stage. Welcome for any pull request if you find something you want but not in this repo yet.

Currently you can do
- like/unlike media
- get timeline feed
- get ranked medias / recent medias of location
- get inbox message
- broadcast to any inbox thread
- approve pending inbox thread

## Example

```go
ig := instagram.Create(username, password) // init an instance

ig.Login() // login Instagram

ig.TimelineFeed.Get() // get timeline feed

ig.Like(ig.TimelineFeed.Items[0].ID) // like the first item of the feed
```

## License

MIT

## Terms and Conditions

- You shouldn't use this repo for spam, massive sending or any malicious activity.
- We don't give support to anyone who wants to violate this terms and conditions.
- We reserve the right to block any user who doesn't meet the conditions.

## Similar Projects

- [instagram-private-api](https://github.com/huttarichard/instagram-private-api) (Node.js)
- [Instagram-API](https://github.com/mgp25/Instagram-API) (PHP)