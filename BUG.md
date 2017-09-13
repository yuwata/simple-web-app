# Publishing port does not work on the second run
```
$ docker run -p 8080:8080 --name hoge app-hello
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
{"message":"Hello World!!"}
$ docker stop hoge
$ docker start hoge
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
curl: (56) Recv failure: Connection reset by peer
```
Two workarounds
1. Use `host` network
1. Restart docker.service
