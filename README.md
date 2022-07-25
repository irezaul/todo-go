# todo-go
A simply activity


### Couchbase install -- >
## On Ubuntu 22.04;
#### First run this command
```bash
echo 'deb [ arch=amd64 ] http://packages.couchbase.com/releases/couchbase-server/community/deb/ focal focal/main' \
> /etc/apt/sources.list.d/couchbase.list
```
* install repository signing key :
```bash
wget -qO- http://packages.couchbase.com/ubuntu/couchbase.key | gpg --dearmor > /etc/apt/trusted.gpg.d/couchbase.gpg
```
* Run system update
```bash
 sudo apt-get update

 apt update
```
* install Couchbase on Ubuntu 22.04
```bash
apt install couchbase-server-community
```
###  Now check couchbase install or not open your browser and type `localhost:8091`

![couchebase-install-done](https://user-images.githubusercontent.com/77927449/180840578-73a9977a-82e9-48ee-a343-5bccf9c926c8.png)
