# edwin
edge service governance base on eBPF.





Note：

加载程序时如果出现

bpftool is now running in libbpf strict mode and has more stringent requirements about BPF programs.

需要加上--legacy，使用legacy模式

```bash
bpftool --legacy prog load bpf_sockops.o /sys/fs/bpf/bpf_sockops type sockops
```

