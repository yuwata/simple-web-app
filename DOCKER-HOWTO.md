* Preparation
  * DNS
If the host system uses systemd-resolved and cannot access google DNS servers,
then you should edit /etc/sysconfig/docker-network like
```
# /etc/sysconfig/docker-network
DOCKER_NETWORK_OPTIONS=--dns 10.224.254.1
```
  * Run container without root privileges
Create group `docker` and belong users to the group.
```
sudo groupadd docker
sudo usermod -aG docker $USER
```

* Build container
```
docker build --rm -t app-hello .
```
If we get an error like "Could not resolve host: github.com", then
try to restart docker.service.

* Start
```
docker run --rm -ti -p 8080:8080 --net=host --name hoge app-hello
```
Why '--net=host' is necessary??

#TODO
- with firewalld
