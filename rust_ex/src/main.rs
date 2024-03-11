use redbpf_probes::xdp::prelude::*;
use redbpf_probes::xdp::XdpContext;
use std::sync::atomic::{AtomicU64, Ordering};

// AtomicU64를 사용하여 패킷 수를 카운트하는 변수를 정의합니다.
static PACKET_COUNT: AtomicU64 = AtomicU64::new(0);

// XDP 프로그램을 정의합니다.
#[xdp]
pub fn count_packets(ctx: XdpContext) -> XdpResult {
    // 패킷이 들어오면 패킷 수를 증가시킵니다.
    PACKET_COUNT.fetch_add(1, Ordering::Relaxed);

    // XDP_PASS를 반환하여 패킷을 기존의 프로세스로 전달합니다.
    XdpAction::Pass
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    // eBPF 프로그램을 로드하고 인터페이스에 연결합니다.
    let mut _socket = count_packets.load()?;

    // 패킷 수를 출력합니다.
    println!("Packet count: {}", PACKET_COUNT.load(Ordering::Relaxed));

    Ok(())
}
