#### FS under Docker:

**AUFS** - low performance
**Overlay2** - tool from linux kernel, fast

[description](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/filesystems/overlayfs.txt)
[description from Docker](https://docs.docker.com/storage/storagedriver/overlayfs-driver/)

**Brief**
Upper/lower layers exist.
*upper/top*: for files with same names upper is in priority, directories merge, writable layer
*lower*: readable, may be other overlay, all scenarious except read are configured by flags
[reads/writes in Docker container/image](https://docs.docker.com/storage/storagedriver/overlayfs-driver/#how-container-reads-and-writes-work-with-overlay-or-overlay2)
folders */upper* + */lower* -> */merged*, really stores in /workdir
suppport up to 128 lower OverlayFS layers

##### Important 
**copy_up** - files changed in image(lowlevel) by container(toplevel) *COPY* to container, even if small piece was changed.
**rename** - not fully supported.


**for examples of overlay2 commands and results check notes/overlay2_examples.md**

**SquashFS** - creates archive with fast read(need CPU, decompression), more expensive to change. People use for USB/live file systems



#### Docker installation
```bash
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

sudo apt-get update

sudo apt-get install docker-ce
sudo docker run hello-world

```

#### Docker Uninstall
```bash
sudo apt-get purge docker-ce
sudo rm -rf /var/lib/docker
```