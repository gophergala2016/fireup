# fireup

Fireup is a tool for dynamically broadcast markdown slides to all browser clients connected to it. It has two modes of operation, the 
server mode and the client mode.


# Installation

```
$ go get github.com/gophergala2016/fireup
```

this will download `fireup` in your `$GOPATH`. Now use `go build`
command within the fireup directory in GOPATH to build the fireup binary. Now place the binary somewhere in your path.

# Usage

Run fireup in server mode with the command 

``
./fireup s
2016/01/25 04:56:08 starting http listener on http://0.0.0.0:8080
2016/01/25 04:56:08 starting push listener on 0.0.0.0:8081

``

and share the url of the http listener to all participants.

now push slides to the server usong the push command

```
$ fireup p ./slides/first.md
```

some cool used of fire up

```
$ head -10 main.go | fireup p
$ for s in `ls samples/*.md` ; do cat $s| fireup p; sleep 4; done
$ echo "<pre>`ls -l`</pre>"| fireup p

```
for a detailed help on the commands and their flags, see the 
usage message below.

```
$ fireup 
NAME:
   fireup - push markdown slides on the fly.

USAGE:
   fireup [global options] command [command options] [arguments...]
   
VERSION:
   0.0.0
   
COMMANDS:
   serve, s	start in server mode
   push, p	push slide.
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version


```

# Starting the server

# Pushing files to the server

# License



