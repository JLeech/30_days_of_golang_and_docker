
```uname -a```
Linux xxx 4.4.0-137-generic #163-Ubuntu SMP Mon Sep 24 13:14:43 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
```df -T (T - show file system type)```
/dev/sda1      ext4

[commands](https://askubuntu.com/questions/699565/example-overlayfs-usage)
```
cd /tmp
mkdir lower upper workdir overlay
sudo mount -t overlay -o \
lowerdir=/tmp/lower,\
upperdir=/tmp/upper,\
workdir=/tmp/workdir \
none /tmp/overlay
```
-t   : set type of mounting (overlay)
none : name of mounted file system
```
echo "upper" > ./upper/from_up
echo "overlay" > ./overlay/from_overlay
echo "lower" > ./lower/from_low
```
```
ls lower/ upper/ overlay/
```
*lower/:*
from_low
*overlay/:*
from_low  from_overlay  from_up
*upper/:*
from_overlay  from_up
```
echo "updated_low" > ./overlay/from_low
ls lower/ upper/ overlay/
```
*lower/:*
from_low
*overlay/:*
from_low  from_overlay  from_up
*upper/:*
from_low  from_overlay  from_up
```
head lower/* upper/* overlay/*
```
*==> lower/from_low <==*
lower

*==> upper/from_low <==*
lower
updated_low

*==> upper/from_overlay <==*
overlay

*==> upper/from_up <==*
upper

*==> overlay/from_low <==*
lower
updated_low

*==> overlay/from_overlay <==*
overlay

*==> overlay/from_up <==*
upper

```
sudo umount -A none
rm -rf /tmp/lower /tmp/upper /tmp/workdir /tmp/overlay
```