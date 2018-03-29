# Dockerbox

## 概述

配合使用docker作为开发环境的工具

## 部署

请先安装[dep](https://golang.github.io/dep/)

```
git clone git@github.com:ligang1109/dockerbox.git
cd dockerbox
./deploy/deploy.sh host1 host2 ...
```

这会将dbox工具安装至目标主机的/usr/local/bin下，并生成常用的dconf.json放到$HOME/.dconf.json

## dconf.json

所有工具都依赖于一个全局配置文件dconf.json，它的默认路径放在$HOME/.dconf.json，也可以手动在执行时通过flag参数指定：-dconfPath=dconfPath

### 结构解析

示例：
```
{
    "nginx": {                            //目标key
        "container_name": "nginx-1.8.1",  //目标容器名
        "exec": {                         //dbox exec相关
            "shell_cmd": "bash -c",       //由于支持执行多个命令，所以这里配置使用容器中的哪个shellCmd执行，如/bin/sh -c
            "pre_cmd": "source ~/.bashrc" //默认先执行哪个命令，例如source ~/.bashrc可以导出自定义环境变量
        }
    },
    "php": {
        "container_name": "php-7.1.9",
        "exec": {
            "shell_cmd": "bash -c",
            "cwd": true,                   //执行命令前先cd到执行dbox命令的目录，下面有示例说明
            "pre_cmd": "source ~/.bashrc"
        }
    }
}
```

## 使用方法

命令格式：

```
dbox [options] command containerKey args ...
```

当前支持options：

- -logLevel: 可选，值见[logLevel](https://github.com/goinbox/golog/blob/master/base.go)

- -dconfPath：可选，放置dconf配置文件的路径，请用绝对路径，默认为：`$HOME/.dconf.json`

其他参数：

- command：必选，dbox支持的命令名称

- containerKey：必选，dconf中的容器key，`all`为全部key

- args：可选，指定cmd的额外args

## command说明

### exec

用途：从容器外部执行容器内的命令

示例运行：

```
ligang@vm-centos7 ~ $ dbox exec nginx nginx -v
nginx version: nginx/1.8.1
```

还可以直接进入容器中的交互命令行，如：

```
ligang@vm-centos7 ~ $ dbox exec mysql mysql -uroot -p123
mysql: [Warning] Using a password on the command line interface can be insecure.
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 5
Server version: 5.7.9-log Source distribution

Copyright (c) 2000, 2015, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 
```

这里用一个实际运用场景来说明下cwd这个配置：

我开发的php程序放在一个路径下：

```
ligang@vm-centos7 ~/tmp/php $ ls
test.php
ligang@vm-centos7 ~/tmp/php $ pwd
/home/ligang/tmp/php
```

如果我想从容器外面调试这个test.php代码，要使用命令：

```
ligang@vm-centos7 ~/tmp/php $ dbox exec php php $PWD/test.php 
hello, world
```

当配置了cwd这一项后，就可以省掉$PWD这个路径了：

```
ligang@vm-centos7 ~/tmp/php $ dbox exec php php test.php 
hello, world
```

容器中实际执行了：

```
ligang@vm-centos7 ~/tmp/php $ dbox -logLevel=1 exec php php test.php 
[debug] [2018-03-26 12:21:10]   cmd: docker exec -it php-7.1.9 bash -c 'cd /home/ligang/tmp/php;source ~/.bashrc;php test.php'
```

是不是方便了很多！

### attach

需要先安装[nsenter](http://man7.org/linux/man-pages/man1/nsenter.1.html)

用途：进入指定容器，多终端不会相互影响

示例:
```
ligang@vm-centos7 ~ $ dbox attach nginx
[sudo] password for ligang: 
[root@vm-centos7 /]# ps auxww
USER        PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root          1  0.0  0.0  11640  1332 ?        Ss   17:10   0:00 /bin/bash /build/script/init.sh
root          5  0.0  0.0  31352  3012 ?        S    17:10   0:00 nginx: master process /usr/local/nginx/sbin/nginx -g daemon off;
nobody        6  0.0  0.1  34756  4656 ?        S    17:10   0:00 nginx: worker process
root        116  0.0  0.0  11880  1904 ?        S    17:50   0:00 -bash
root        129  0.0  0.0  47560  1692 ?        R+   17:50   0:00 ps auxww
```

### start

用途：启动容器

示例:
```
ligang@vm-centos7 ~ $ dbox start nginx
[warning]	[2018-03-15 17:53:15]	start container: nginx-1.8.1
nginx-1.8.1
```

配置容器开机启动：
```
ligang@vm-centos7 ~ $ cat /etc/systemd/system/dstart.service 
[Unit]
Description=Docker Container Starter
After=network.target docker.service
Requires=docker.service

[Service]
Type=forking
ExecStart=/usr/local/bin/dbox -dconfPath=/home/ligang/.dconf.json start all

[Install]
WantedBy=multi-user.target
```

### stop

用途：停止容器

示例:
```
ligang@vm-centos7 ~ $ dbox stop nginx
[warning]	[2018-03-15 17:51:07]	stop container: nginx-1.8.1
nginx-1.8.1
```

### restart

用途：重启容器（先stop，再start）

示例:
```
ligang@vm-centos7 ~ $ dbox restart nginx
[notice]	[2018-03-15 17:54:33]	restart container: nginx
[warning]	[2018-03-15 17:54:33]	stop container: nginx-1.8.1
nginx-1.8.1
[warning]	[2018-03-15 17:54:43]	start container: nginx-1.8.1
nginx-1.8.1
```

### rm

用途：删除容器

示例:
```
ligang@vm-centos7 ~ $ dbox rm nginx
[warning]	[2018-03-29 16:43:43]	rm container: nginx-1.8.1
nginx-1.8.1
```

### rmi

用途：删除容器及容器使用的镜像

示例:
```
ligang@vm-centos7 ~ $ dbox rmi nginx
[warning]	[2018-03-29 16:46:30]	rm container: nginx-1.8.1
nginx-1.8.1
[warning]	[2018-03-29 16:46:30]	rmi image: nginx-1.8.1
Untagged: andals/nginx:1.8.1
Deleted: sha256:ccc9c3c7f1a54910964b05bd68c9b8d88847dbcbdb9da203ba24ae7630f5f00d
Deleted: sha256:a7df818e38f0a6cabfb939b634ad57914a27393fd5f6f13d435a521802a3da25
Deleted: sha256:23e59bec3f189dcbfc4d8cc1cbaf75b829f9bbe26ffbfd799c4ed32e9051f794
Deleted: sha256:b768a6f4cc91e7409d663b8a2f97e361bbe76210bc0d6ed3519319611ecab9f6
Deleted: sha256:03c749ca45f33d140deffa43e24663a8ed6858fdb00704fd23f4e758221b4221
Deleted: sha256:fcb5068ea905daaaafc83b9bbce878e0222c2657fb9e1b983bad342534604a1d
Deleted: sha256:1ff0f8993b098be3071ae9b2c61caad17f2ce501765f195c0accaf47dd32793b
Deleted: sha256:33c294a729df0e7dfa538a348b3b4d31cc626e75c2b8bf099b6642eb05debcc7
Deleted: sha256:7dcecd60fc0bb263dc0b38d14c3a367cea6b0a8e91deb4bcbd31256050db7ea0
```

#### 特别说明

使用start、stop、restart、rm、rmi这几个命令时，如果containerKey为`all`，则取dconf中的全部key执行操作
