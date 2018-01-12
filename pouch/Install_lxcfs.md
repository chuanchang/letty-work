# Install lxcfs

# On Ubuntu
```
apt-get install lxcfs
```

# On CentOS

You'll have to build from souce code on CentOS to install lxcfs. [Here](https://linuxcontainers.org/lxcfs/getting-started/) is the official instruction about building and running lxcfs from git. Here are the steps:
	
1. Install autotools 

	```
	# yum install autotools-dev m4 autoconf2.13 autobook autoconf-archive gnu-standards autoconf-doc libtool
	```
	
2. Install dependent libs

	```
	# yes | yum install fuse-devel.x86_64
	# yes | yum install pam-devel.x86_64
	```
	
3. Build from source code

	```
	git clone git://github.com/lxc/lxcfs
	cd lxcfs
	./bootstrap.sh
	./configure
	```
	Please make sure the `configure` succeeds. Or you may need to solve the failure before going to the next step. The build the source code.
	
	```
	# make install
	# mkdir -p /var/lib/lxcfs
	# ./lxcfs -s -f -o allow_other /var/lib/lxcfs/
	```
	