# http-gateway - 统一http代理

http-gateway扮演着一个反向代理的角色,但是他可以同时代理多个后台服务,并且可以通过管理页面动态增加后台服务,每个后台应用对应于proxy的子域名

例如

	http://hello.proxy.yuankui.me -> http://192.168.1.1:8081
	http://kitty.proxy.yuankui.me -> http://192.168.2.1:8083

因此该项目包括两部分

- admin页面 - 用于管理后台应用
- proxy - 用于后台服务的反向代理

## requirements

- golang
- mysql-server
- [godep](https://github.com/tools/godep/)

## quick start

### 1.get the source

	go get github.com/yuankui/http-gateway
	cd $GOPATH/src/github.com/yuankui/http-gateway
	
	godep restore

### 2. bind dns

	*.proxy.your.domain -> 127.0.0.1

### 3. create database

	mysql -uroot	
	create database test_proxy;
	
### 4. modify config

	vi conf/app.conf	

### 5. start server
	
	
	go build .
	
	## admin
	./http-server
	
	## proxy
	cd proxy-server
	go build .
	./proxy-server
	
	
- visit `http://localhost:8081` to add backend app
- visit `http://app_name.proxy.your.domain` to visit the proxied backend app