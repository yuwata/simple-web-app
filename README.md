# Ex.1. Create a simple web app with gin and docker

The source of the application is located at `src/hello/hello.go`,
and managed by [gb](https://getgb.io/). So, you can test the code with
```
$ gb vendor fetch github.com/gin-gonic/gin
$ gb build all
$ ./bin/hello
```
If `PORT` env is not set, the application is served on `:8080`.

* Build container
```
$ docker build -t hello -f docker/hello/Dockerfile .
```
* Start container
```
$ docker run -p 8080:8080 --name hoge hello
```
* Connect container
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
{"message":"Hello World!!"}
```

# Ex.2. Deploy the previous app with a PaaS/IaaS

I use [Heroku](https://www.heroku.com/) for publishing the app. The URL is https://vast-hamlet-24650.herokuapp.com/.
```
$ curl -XGET -H 'Content-Type:application/json' https://vast-hamlet-24650.herokuapp.com/
{"message":"Hello World!!"}
```

# Ex.3 Create a RESTful application

* With docker-compose
```
$ docker-compose up
```

* Without docker-compose

If you want to test the app without `docker-compose`, then
```
$ docker pull postgres:latest
$ docker build -t restful -f docker/restful/Dockerfile .
$ docker run --rm --name db -e POSTGRES_PASSWORD=foo -d postgres
$ docker run --rm -ti -p 8080:8080 --name foo -e DATABASE_URL="host=db user=postgres password=foo sslmode=disable" --link db:db restful
```

* Heroku

Also, the application is running on Heroku with `heroku-postgresql` addon.
The URL is https://vast-hamlet-24650.herokuapp.com/.
