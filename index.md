# Dockerbox

## 概述

docker环境开发的配合工具

## 部署

请先安装[dep](https://golang.github.io/dep/)

```
git clone git@github.com:ligang1109/dockerbox.git
cd dockerbox
./deploy/deploy.sh host1 host2 ...
```

这会将相关工具安装至目标主机的/usr/local/bin下，并生成常用的dconf.json放到$HOME/.dconf.json

### dconf.json

###### 概述

所有工具都依赖于一个全局配置文件dconf.json，它的默认路径放在$HOME/.dconf.json，也可以手动在执行时通过flag参数指定：-dconf=dconfPath

###### 结构解析

示例：
```
{
    "nginx": {                                      //目标key
        "container_name": "nginx-1.8.0",            //目标容器名
        "exec_path": "/usr/local/nginx/sbin/nginx"  //dexec执行时，实际执行的目标容器中的可执行文件的绝对路径
    },
    "php": {
        "container_name": "php-5.5.27",
        "exec_path": "sudo -u ligang /usr/local/php/bin/php"
    }
}
{
    "nginx": {
        "container_name": "nginx-1.8.1",  //
        "exec": {
            "shell_cmd": "bash -c",
            "pre_cmd": "source ~/.bashrc"
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


#### 工具介绍

###### dattach

用途：进入指定容器，多终端不会相互影响（需要nsenter）

示例:
```
dattach nginx
```

###### dexec

用途：从容器外部执行容器内的命令

示例:
```
dexec nginx -s reload
```

###### dstart

用途：启动容器

示例:
```
dstart nginx
```

配置容器开机启动：
```
ligang@vbox02 ~ $ cat /etc/systemd/system/dstart.service 
[Unit]
Description=Docker Container Starter
After=network.target docker.service
Requires=docker.service

[Service]
Type=forking
ExecStart=/usr/local/bin/dstart -dconf=/home/ligang/.dconf.json all

[Install]
WantedBy=multi-user.target
```
