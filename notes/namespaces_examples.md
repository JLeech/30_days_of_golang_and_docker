### Networking
### NOT OVER


[examples](https://docker-k8s-lab.readthedocs.io/en/latest/docker/netns.html)

check existing nets namespaces
```bash
ip netns list

sudo ip netns add test1
```
netns - manage network namespaces

```bash
sudo ip netns exec test1 ip a
```
check network interfaces inside test1 namespace
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

LOOPBACK - interface for self connection
mtu (maximum transfission unit) 65536 - size of block to transfer without fragmentation
qdisc - type of network scheduler
noop - ? 
state DOWN - turned off

[good article with pictures and commands](http://www.opencloudblog.com/?p=66)