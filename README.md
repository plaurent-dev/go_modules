# GO Modules

This guide has been created using an Ubuntu 18.04, in consequence, folders, commands and results may vary.

There are a lot of stuff regarding modules around, but after reading it, i still had some difficulty understanding how modules work and how i can use them.

Note : I had no golang background at all

Version used is **go1.14.2 linux/amd64**

## Prerequisite

GOPATH has to be properly configured. Mine has been added within **~/.bashrc** even if i think it could be a better idea to use .profile instead.

```
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

```
(base) pierre@pierre-MS-7B45:~/Documents/modules$ which go
/usr/local/go/bin/go
(base) pierre@pierre-MS-7B45:~/Documents/modules$ go version
go version go1.14.2 linux/amd64
(base) pierre@pierre-MS-7B45:~/Documents/modules$ go env GOPATH
/home/pierre/go
```



## GITHUB structure

**module**

​	hello.go

​	hello_test.go

**program**

​	usehellomod.go	

------

usehellomod.go will use two kind of packages,

- an inner one, fmt
- a GITHUB one

hello.go and hello_test.go will be used to create a module stored on your own github.



## MODULE (Create/Push)

### Automatically create a go.mod

```bash
go mod init github.com/username/hello

go mod tidy

go test ./...

git add .

git commit -m "hello: changes for v0.1.0"

git tag v0.1.0

git push origin v0.1.0
```



### Useful

Is v0.1.0 version available ?

```bash
go list -m github.com/plaurent-django/hello@v0.1.0
```

What is hello package composed of ?

```bash
go doc github.com/plaurent-django/hello
```

```bash
package hello // import "github.com/plaurent-django/hello"

func Hello() string
func Person() string
func Proverb() string
```



## CALLING MODULE

```bash
go mod init github.com/plaurent-django/workmodule

go run callmod.go
```

## GET A NEW PACKAGE

```
go get github.com/plaurent-django/hello@v1.0.8
```

## GIT COMMANDS

```
(base) pierre@pierre-MS-7B45:~/Documents/modules$ git tag
v0.1.0
v1.0.0
v1.0.5
v1.0.6
v1.0.7
v1.0.8

(base) pierre@pierre-MS-7B45:~/Documents/modules$ git tag -nv0.1.0          hello: initial commit
v1.0.0          hello: initial commit
v1.0.5          hello: changes for v1.0.5
v1.0.6          hello: changes for v1.0.6
v1.0.7          hello: changes for v1.0.7
v1.0.8          hello: changes for v1.0.8

(base) pierre@pierre-MS-7B45:~/Documents/modules$ git tag -n --sort=committerdate
v0.1.0          hello: initial commit
v1.0.0          hello: initial commit
v1.0.5          hello: changes for v1.0.5
v1.0.6          hello: changes for v1.0.6
v1.0.7          hello: changes for v1.0.7
v1.0.8          hello: changes for v1.0.8

(base) pierre@pierre-MS-7B45:~/Documents/modules$ git ls-remote --tags
From https://github.com/plaurent-django/hello.git
cde0194dcda2ba5c38205e38f999777b76ec57c4	refs/tags/v0.1.0
cde0194dcda2ba5c38205e38f999777b76ec57c4	refs/tags/v1.0.0
97ee502235dd45449e66ccaace142b333a8a72e3	refs/tags/v1.0.5
b9b8b91e819f30846b7ec6654069a5c113cbb61b	refs/tags/v1.0.6
572aaa4368f9a313b21fe896d30431a335f79566	refs/tags/v1.0.7
9262413087a0e163bfd6dce698c7f0650ad6be51	refs/tags/v1.0.8

(base) pierre@pierre-MS-7B45:~/Documents/modules$ git ls-remote --tags origin
cde0194dcda2ba5c38205e38f999777b76ec57c4	refs/tags/v0.1.0
cde0194dcda2ba5c38205e38f999777b76ec57c4	refs/tags/v1.0.0
97ee502235dd45449e66ccaace142b333a8a72e3	refs/tags/v1.0.5
b9b8b91e819f30846b7ec6654069a5c113cbb61b	refs/tags/v1.0.6
572aaa4368f9a313b21fe896d30431a335f79566	refs/tags/v1.0.7
9262413087a0e163bfd6dce698c7f0650ad6be51	refs/tags/v1.0.8

(base) pierre@pierre-MS-7B45:~/Documents/modules$ git fetch --all --tags
Récupération de origin

(base) pierre@pierre-MS-7B45:~/Documents/modules$ git describe
fatal: Aucune étiquette annotée ne peut décrire '9262413087a0e163bfd6dce698c7f0650ad6be51'.
Cependant, il existe des étiquettes non-annotées : essayez avec --tags.
```

## LINKS

[golang.org](https://golang.org/doc/code.html).

https://blog.golang.org/using-go-modules
