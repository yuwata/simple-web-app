# Preparation
* DNS

If the host system uses systemd-resolved and cannot access google DNS servers,
then you should edit /etc/sysconfig/docker-network like
```
# /etc/sysconfig/docker-network
DOCKER_NETWORK_OPTIONS=--dns 10.224.254.1
```
* Run container without root privileges

Create group `docker` and add users to the group.
```
sudo groupadd docker
sudo usermod -aG docker $USER
```
You need to close the session to update the user info.

# Build container
```
docker build --rm -t hello -f docker/hello/Dockerfile .
```
If we get an error like "Could not resolve host: github.com", then
try to restart docker.service. This may be related to the bug listed in BUG.md.

# Start
```
docker run --rm -ti -p 8080:8080 --name hoge hello
```
If we cannot connect localhost:8080, then try to use `--network host` option
or restart docker.service. See BUG.md for detail.

# OPTIONAL
```
docker pull postgres:latest
docker build --rm -t restful -f docker/restful/Dockerfile .
docker run --rm --name foo -e POSTGRES_PASSWORD=foo -d --network host postgres
docker run --rm -ti -p 8080:8080 --name bar -e DATABASE_URL="host=localhost user=postgres password=foo sslmode=disable" --network host restful
```

# TODO
- with firewalld?
