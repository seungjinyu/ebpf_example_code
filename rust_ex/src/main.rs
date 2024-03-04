use redbpf_probes::kprobe::prelude::*;

program!(0xFFFFFFFE, "GPL");

#[kprobe("__tcp4_proc_connect")]
pub fn handle_connect(ctx: BPFContext) {
    bpf_printk!("Hello World! Someone connected via SSH\n");
}
