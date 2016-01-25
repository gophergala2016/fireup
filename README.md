# fireup

Fireup is a tool for dynamically broadcasting markdown slides to all browser clients connected to it. It has two modes of operation, the 
server mode and the client mode. 

Fireup can be thought of as a virtual whiteboard for a conference or any such collaborative endeavour. Start `fireup` in server mode
and share it's web url to all participants or on a projector. Now all participants can push their content to the server. 

# Installation
To install `fireup` use the following command

(**NOTE:** you must have [Go](http://golang.org) installed before you can install fireup)

```
$ go get github.com/gophergala2016/fireup
```

this will download `fireup` in your `$GOPATH`. Now use `go build`
command within the fireup directory in GOPATH to build the fireup binary. Now place the binary somewhere in your path.

# Usage

Run fireup in server mode with the command 

```
$ fireup s
2016/01/25 04:56:08 starting http listener on http://0.0.0.0:8080
2016/01/25 04:56:08 starting push listener on 0.0.0.0:8081

```

and share the url of the http listener to all participants. At this point all browsers pointing to
this url will see the following screen

![](https://raw.githubusercontent.com/gophergala2016/fireup/master/screenshots/empty.png)


Now push slides to the server using the push command as follows

```
$ fireup p ./samples/gopher.md
```
fireup then broadcasts the slide to all browser clients connected to fireup instance running in server mode.
The screenshot below displays the slide pushed to the browser window.

![](https://raw.githubusercontent.com/gophergala2016/fireup/master/screenshots/gopher.png)

some cool uses of fireup

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

## Sample slides

For some sample slides, look in the [samples](https://github.com/gophergala2016/fireup/tree/master/samples) folder.

## Features to be added soon

1. Support for file types other than markdown.
2. Support for pushing audio and video files.



# License

MIT License. [See here](https://github.com/gophergala2016/fireup/blob/master/LICENSE)

