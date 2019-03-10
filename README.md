bashserver is a simple appserver that runs a shell command. THe shell command is passed as a parameter. Each request runs the command and returns the result

Bashserver listens on port 8080 and has two endpoints

| Path     | Action                                         |
|----------|------------------------------------------------|
| /        | Run the supplied command and return its output |
| /version | Return the verision of the app (the contents of the $VERSION environemnt variable |


Example usage

```
docker run -p8080:8080 bashserver ls
```

Call the server

```
$ curl localhost:8080
bin
dev
etc
home
lib
media
mnt
opt
proc
root
run
sbin
srv
sys
tmp
usr
var
```



