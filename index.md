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

#### 特别说明

使用start、stop、restart这几个命令时，如果containerKey为`all`，则取dconf中的全部key执行操作