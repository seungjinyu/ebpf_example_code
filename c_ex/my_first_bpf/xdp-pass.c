#include "common.h"

#define XDP_PASS 2

SEC("xdp")
// int xdp_prog_simple(struct xdp_md *ctx)
int xdp_prog_simple()
{
    return XDP_PASS;
}


char _license[] SEC("license") ="GPL";