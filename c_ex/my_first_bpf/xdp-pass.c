//go:build ignore


#include "common.h"

#define __ctx_buff xdp_md
#define ETH_ALEN	6

struct  {
    __uint (type , BPF_MAP_TYPE_HASH);
    __type(key, __u32);
    __type(value, __u64);
    __uint(max_entries, 1);
} my_map SEC(".maps");


#define XDP_PASS 2

SEC("xdp")
int xdp_prog_simple(struct xdp_md *ctx)
{
    __u32 data_meta_value = ctx->data;
    __u32 key = 0;

    

    bpf_printk("packet recedved");

    // __u64 *valp;

    // valp = bpf_map_lookup_elem(&my_map, &key);

    // if (!valp){
    //     bpf_map_update_elem(&my_map,&key, &data_meta_value, BPF_ANY);
    //     return 0;
    // }
    
    // bpf_map_update_elem(&my_map, &key, &data_meta_value, BPF_ANY);

    return XDP_PASS;
}


char _license[] SEC("license") ="GPL";