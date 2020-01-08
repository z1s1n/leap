# Go - based local board jumper

![mascot](mascot.png)

The terminal software most used on `MAC` is `item2`, but `item2` is inconvenient to connect to the server (remember the password). To configure `profile`, if you want to connect to multiple machines, you need to configure multiple `profiles`. If you want to add personalized Settings: background blur, font, etc., you need to modify each `profile`, which is very inconvenient.

This project through the establishment of the principle of the board jumper, support the selection of serial number to log in the machine.

## support:
- [x] login by username and password
- [x] login by username and key (no password and with password)
- [x] serial number selection
- [x] search

## install:

1. `go get github.com/chazyu1996/leap`
2. `cd $GOPATH/src/github.com/chazyu1996/leap`
3. `go build leap.go`
4. Edit the `config.yaml` file

## screenshot:

![screenshot](screenshot.png)
