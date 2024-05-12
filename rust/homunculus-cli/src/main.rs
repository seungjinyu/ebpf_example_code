use clap::Parser;
use rusoto_core::Region;
use rusoto_credential::DefaultCredentialsProvider;
use rusoto_eks::{DescribeClusterInput, Eks, EksClient};



#[derive(Parser, Debug)]
#[command(version, about, long_about = None)]
struct Args {

    #[arg(short, long)]
    name: String,

    #[arg(short , long, default_value_t = 1 )]
    count: u8,

}

fn main() {

    println!("Hello, world!");

    let args = Args::parse();

    for _ in 0..args.count{
        println!("Hello {}", args.name);
    }
    // AWS 리전 및 자격 증명 제공자 설정
    let region = Region::default();
    let credentials_provider = DefaultCredentialsProvider::new().unwrap();

    // Amazon EKS 클라이언트 생성
    let client = EksClient::new_with(credentials_provider, region);

    // Amazon EKS 클러스터 이름 지정
    let cluster_name = "your-cluster-name";

    // DescribeClusterInput 구성
    let input = DescribeClusterInput {
        name: cluster_name.to_string(),
        ..Default::default()
    };

    // Amazon EKS 클러스터 정보 조회
    match client.describe_cluster(input).await {
        Ok(output) => {
            if let Some(cluster) = output.cluster {
                println!("Cluster name: {}", cluster.name.unwrap_or_default());
                println!("Cluster status: {:?}", cluster.status.unwrap_or_default());
                println!("Cluster version: {:?}", cluster.version.unwrap_or_default());
                // 기타 클러스터 정보 출력
            } else {
                println!("Cluster not found.");
            }
        }
        Err(e) => {
            eprintln!("Error describing cluster: {}", e);
        }
    }

    
}
