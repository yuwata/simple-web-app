* install heroku toolkit
* `heroku login`
* run `heroku create` in the project dir
* install gb by `go get github.com/constabulary/gb/...`
* create project tree like
```
└── src
      └── hello
            └── hello.go
```
* set PATH env as `gb` and `gb-vendor` can run.
* fetch libraries by `gb vendor fetch github.com/gin-gonic/gin`
* try to build by `gb build all`
* create Procfile as
```
web: hello
```
* deploy the app by `git push heroku master`
* see logs by `heroku logs --tail`
* try to connect by `heroku open` or
```
curl -XGET -H 'Content-Type:application/json' https://vast-hamlet-24650.herokuapp.com/
```

# Tips
* If PATH env is not correctly set, then `gb` command cannot find `gb-vendor`. Thus, you cannot run `gb vendor` command.
* heroku set PORT env and the port is mapped to :80. So, in the application, you should use PORT env as the service port (gin automatically uses PORT env).
