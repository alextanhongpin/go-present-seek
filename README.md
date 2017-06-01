# golang workshop

Link to [presentation](https://go-talks.appspot.com/github.com/alextanhongpin/go-present/main.slide#1)

- Work in pairs, learn the basics of golang and build a web api
- Packaging, building the app
- Syntax and best practices

## Tools required for the workshop

- mysql - https://dev.mysql.com/downloads/mysql/
- golang - https://nats.io/documentation/tutorials/go-install/
- an text editor - SublimeText, VisualStudioCode or Atom will do.


## IDE Setup

visual studio - https://github.com/Microsoft/vscode-go
sublime text 3 - https://github.com/DisposaBoy/GoSublime
atom - https://atom.io/packages/go-plus

## Mysql installation

After installing mysql, you should be able to invoke the mysql command from the terminal. If it doesn't work, try pasting this into your .bash_profiel

```bash
$ PATH=${PATH}:/usr/local/mysql/bin
```

The next thing to do is to reset the password.

```bash
$ mysql -u root - p
# Enter the password that is provided by oracle
$  SET PASSWORD FOR 'root'@'localhost' = PASSWORD('MyNewPass');
```


### List of useful mysql commands
http://g2pc1.bu.edu/~qzpeng/manual/MySQL%20Commands.htm
https://www.pantz.org/software/mysql/mysqlcommands.html



## Useful go packages

```bash
$ go get -u golang.org/x/tools/cmd/goimports
$ go get -u golang.org/x/tools/cmd/vet
$ go get -u golang.org/x/tools/cmd/oracle
$ go get -u golang.org/x/tools/cmd/godoc
```

Command goimports updates your Go import lines, adding missing ones and removing unreferenced ones.
```bash
$ go build
$ go run main.go
$ go test
$ go benchmark
```