# Unikorn - Documentation

Welcome to the documentation for unikorn! You will learn a lot of thing about unikorn here.

## Links

- [Commands](https://github.com/5elenay/unikorn/blob/main/docs/Commands.md)
- [Folder Structure](https://github.com/5elenay/unikorn/blob/main/docs/FolderStructure.md)
- [Making Your Own Unikorn Package](https://github.com/5elenay/unikorn/blob/main/docs/CreatingPackages.md)

## Installation

Check [release page](https://github.com/5elenay/unikorn/releases/latest) for binary files. Just download the file for your operating system and check with:

- **Windows**: `.\unikorn.exe version`
- **Linux/macOs**: `./unikorn version`

Now you only need to add unikorn to the path and you are done!

## Using a Unikorn Package

Since all of the packages in the `unikorn` folder, we can do this:

```py
from unikorn import helloworld

helloworld.hello() # hello world
```

## Compiling

If you can't find a release for your system, you can compile Unikorn yourself easily.

- **1 -** Using Compile Script
- **2 -** Compiling Yourself

### Pre-Requests

Make sure you have Go installed. Check with `go version`. If you don't have Go installed, Download latest version from [official website](https://golang.org/dl/)

#### 1-) Using Compile Script

For GNU/Linux, Windows and MacOs you can use compile script without any Go knowledge.

##### Windows

Run this commands in unikorn folder:

```
.\compile\win.bat
```

And it will compile itself.

##### GNU/Linux or MacOs

Run `./compile/linux-macos.sh` And it will compile itself.

#### 2-) Compiling Yourself

#### Step by Step:

- Go to the `/src/` folder.
- Open terminal in `/src/` folder.
- Type `go build`
- Rename the output file with `unikorn`
- Check if it works with `./unikorn help`
