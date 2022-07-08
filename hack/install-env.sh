#!/bin/bash

set x

# centOs use yum/ ubuntu use apt-get
yum install -y make clang llvm libelf-dev libbpf-dev bpfcc-tools \
libbpfcc-dev linux-tools-$(uname -r) linux-headers-$(uname -r)

yum install -y make gcc libssl-dev bc libcap-dev \
  clang gcc-multilib llvm libncurses5-dev git pkg-config libmnl-dev bison flex \
  graphviz build-essential strace tar libfl-dev libedit-dev zlib1g-dev
