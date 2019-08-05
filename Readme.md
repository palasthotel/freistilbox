# freistilbox

freistilbox is a command line tool which helps administering
freistil sites.

This Readme will receive more love and content as the project
progresses. For now, there are just two commands implemented
and they're not that helpful.

The Freistilbox API documentation can be found [here](https://dashboard.freistilbox.com/docs/apis/v1)

## structure

the command package defines an interface every command has to
follow. The main package builds up a list of all commands.

## lazy man starting point

* install go
* `$ mkdir ~/golang`
* `$ export GOPATH=~/golang`
* `$ export PATH=$PATH:~/golang/bin`
* `$ go get github.com/palasthotel/freistilbox`

## using freistilbox

first things first: freistilbox needs your credentials.
So you should start with `freistilbox help` and `freistilbox credentials USER TOKEN`
.
Afterwards, you can use all available commands.
