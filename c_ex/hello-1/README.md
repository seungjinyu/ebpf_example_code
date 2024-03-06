# basic example for kernel programming 

# create, compile, install modules 

make 

sudo insmod hello-1.ko

sudo lsmod | grep hello

sudo rmmod hello_1
