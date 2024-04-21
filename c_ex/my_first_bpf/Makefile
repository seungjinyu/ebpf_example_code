
xdp-pass.ll: 
	clang -I./ -D__NR_CPUS__=4 -O2 -g --target=bpf -std=gnu89 -nostdinc -emit-llvm -Wall -Wextra -Werror -Wshadow -Wno-address-of-packed-member -Wno-unknown-warning-option -Wno-gnu-variable-sized-type-not-at-end -Wdeclaration-after-statement -Wimplicit-int-conversion -Wenum-conversion -Wimplicit-fallthrough -c xdp-pass.c -o xdp-pass.ll

xdp-pass.o:
	llc -march=bpf -mcpu=v3 -filetype=obj -o xdp-pass.o xdp-pass.ll
