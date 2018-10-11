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


#### Overview

Client <-Sockets/Hetwork-> Daemon

Daemon manage images, containers, network and volumes
Daemon <-> Daemon
Client <-> Daemon_1
	   <-> Daemon_2

user <-> Client <-> Daemon

**Docker registry** - store Docker images. can be private.
**Dockerfile** - is used to customise images. Each command create separate layer(up to 128?) DETAILS?
**Container** - is runnable instance of image.
**Services** - tool to manage many containers. DETAILS?

##### Technologies
**[Namespaces](https://en.wikipedia.org/wiki/Linux_namespaces)** : container like a lone user of system.

*pid*: (PID: Process ID) - ID inside namespaces uniq, but can intersect between namespaces. 
*net*: (NET: Networking). manage virtual interfaces and their connections
*ipc*: Managing access to IPC resources (IPC: InterProcess Communication). **READ**
*mnt*: Managing filesystem mount points (MNT: Mount). 
*uts*: Isolating kernel and version identifiers. (UTS: Unix Timesharing System). **READ**

**Control** : 
**UnionFS** : see LayerFS
**Container** : 


#### DockerFile