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
try to restart docker.service. This may be related to the bug listed in BUG.md.

* Start
```
docker run --rm -ti -p 8080:8080 --name hoge app-hello
```
If we cannot connect localhost:8080, then try to use `--network host` option
or restart docker.service. See BUG.md for detail.

#TODO
- with firewalld
