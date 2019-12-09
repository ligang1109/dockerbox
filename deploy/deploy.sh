#!/bin/bash

curPath=`dirname $0`
cd $curPath/../
prjHome=`pwd`

if [ $# -lt 1 ]
then
    echo "Usage $0 host1 host2 ..."
    exit 1
fi

echo "Make sure you have sshpass in PATH"
echo "Make sure you have sudo permission of online hosts"

echo "Enter username of online hosts: "
read username
echo "Enter your password of online hosts: "
read -s password
sshCmd="sshpass -p $password ssh -o StrictHostKeyChecking=no"
scpCmd="sshpass -p $password scp -o StrictHostKeyChecking=no"


echo "Building binary"
deployTmpDir=$prjHome/tmp/deploy
if [ -d $deployTmpDir ]
then
    rm -rf $deployTmpDir
fi
mkdir -p $deployTmpDir

binName=dbox
cd $prjHome/dbox/main/dbox
go build -o ${binName}.bin main.go 
mv ${binName}.bin $deployTmpDir/$binName

installDstDir=/usr/local/bin
cd $deployTmpDir
for host in $*
do
    echo Deploy to $host
    $scpCmd $binName $username@$host:./
    $sshCmd -t $username@$host "echo $password | sudo -S mv $binName $installDstDir"
#    $scpCmd $prjHome/dconf.json.demo $username@$host:./.dconf.json
done

rm -rf $deployTmpDir
