# make xdp-pass.ll , xdp-pass.o 

make xdp-pass.ll

make xdp-pass.o

# attach the ebpf code on xdp

ip link set dev eth0 xdp obj xdp-pass.o section xdp

# detach the module on xdp

ip link set dev eth0 xdp off

# llvm install 
apt-get install llvm

# ftrace bpf_printk
mount -t debugfs none /sys/kernel/debug

# ping to the right ip