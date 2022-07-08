#!/bin/bash

set -x
# 导入 ELRepo 仓库的公共密钥
rpm --import https://www.elrepo.org/RPM-GPG-KEY-elrepo.org

# 安装 ELRepo 仓库的 yum 源
yum install -y https://www.elrepo.org/elrepo-release-7.el7.elrepo.noarch.rpm

# 替换为清华 ELRepo 源
sed -i "s/mirrorlist=/#mirrorlist=/g" /etc/yum.repos.d/elrepo.repo
sed -i "s#elrepo.org/linux#mirrors.tuna.tsinghua.edu.cn/elrepo#g" /etc/yum.repos.d/elrepo.repo

# (可选) 更新 yum 缓存
yum makecache

# 查看可用的内核版本，kernel-ml（mainline stable）：稳定主线版本，kernel-lt（long term support）：长期支持版本
#yum --disablerepo="*" --enablerepo="elrepo-kernel" list available

# 升级为主线版本
yum --enablerepo=elrepo-kernel install kernel-ml -y

# 查看可用内核版本及启动顺序
sudo awk -F\' '$1=="menuentry " {print i++ " : " $2}' /boot/grub2/grub.cfg

# 查看启动顺序
yum install -y grub2-pc
grub2-editenv list

# 设置开机启动
grub2-set-default 0

# 重启生效
# reboot