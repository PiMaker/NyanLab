# NyanLab

NyanLab is a simple web application that allows images to scroll seamlessly over multiple monitors and even multiple computers if they are networked together.

## Usage

Install this program by calling

```
go get -v -u github.com/PiMaker/NyanLab
```

(you will need the Go language installed on your machine)

Next, call `NyanLab` without parameters from your command line.

Access the server at port :8080. Note that the order in which you access the server on different machines determines the path that your image will take.

Hit enter in the server console to begin the madness!

### Different images

To choose a different image, you need to go into `nyanlab.go` and change the line

```go
const IMGSRC = "nyan.png"
```

to read the filename of your image (which can be an animated .gif BTW!):

```go
const IMGSRC = "gandalf.gif"
```

Then, reinstall NyanLab by calling

```
go install
```

and your done! Note that the background color is white, except when using the default nyan cat image. You can manually change it in the code as well.

## Why?

Because.

## License

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
​                     Version 2, December 2004 

Copyright (C) 2004 Sam Hocevar <sam@hocevar.net> 

Everyone is permitted to copy and distribute verbatim or modified 
copies of this license document, and changing it is allowed as long 
as the name is changed. 

DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION 

0. You just DO WHAT THE FUCK YOU WANT TO.