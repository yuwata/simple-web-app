# Preparation
* systemd-networkd

The default settings of systemd-networkd conflicts docker.
So, create the following files:
```
$ cat /etc/systemd/network/00-docker-bridge.network
[Match]
Driver=bridge

[Link]
Unmanaged=yes
$ cat /etc/systemd/network/00-docker-veth.link
[Match]
Driver=veth

[Link]
MACAddressPolicy=random
$ cat /etc/systemd/network/00-docker-veth.network
[Match]
Driver=veth

[Network]
DHCP=no
IPv6AcceptRA=no
```
* DNS

If the host system cannot access google DNS servers,
then you should edit /etc/sysconfig/docker-network like
```
# /etc/sysconfig/docker-network
DOCKER_NETWORK_OPTIONS=--dns 10.224.254.1
```
* Run container without root privileges

Create group `docker` and add users to the group.
```
sudo groupadd -r docker
sudo usermod -aG docker $USER
```
You may need to close the session to update the user info,
and docker.service may need to be restarted.

* TODO
- with firewalld?

# Build container
```
docker build --rm -t hello -f docker/hello/Dockerfile .
```

# Start
```
docker run --rm -ti -p 8080:8080 --name hoge hello
```

# Link to other containers
```
docker pull postgres:latest
docker build --rm -t restful -f docker/restful/Dockerfile .
docker run --rm --name db -e POSTGRES_PASSWORD=foo -d postgres
docker run --rm -ti -p 8080:8080 --name foo -e DATABASE_URL="host=db user=postgres password=foo sslmode=disable" --link db:db restful
```
