# basic example for kernel programming 

# create, compile, install modules 

make 

sudo insmod hello-1.ko

sudo lsmod | grep hello

sudo rmmod hello_1

sudo journalctl --since "1 hour ago" | grep kernel

# module parameter 권한 

sysfs 에 노출 되지 않으려면 권한 비트를 0 으로 설정해야 한다. 

S_IRUSR: 소유자에 대한 읽기 권한 (User Read)
S_IWUSR: 소유자에 대한 쓰기 권한 (User Write)
S_IRGRP: 그룹에 대한 읽기 권한 (Group Read)
S_IWGRP: 그룹에 대한 쓰기 권한 (Group Write)
S_IROTH: 기타 사용자에 대한 읽기 권한 (Others Read)
