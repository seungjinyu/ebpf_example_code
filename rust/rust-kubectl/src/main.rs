use reqwest::{Client, Response};

#[tokio::main]
async fn main() {
    let client = reqwest::Client::new();
    let url = "https://google.com";
    // request 하고 나서 result 값을 받아 온다
    let res = reqwest::get(url).await.expect("request failed");

    println!(
        "Status code {:?} version: {:?}",
        res.status(),
        res.version()
    );
    println!("header is {:?}", res.headers());

    let res_2 = request_login(client, url);

    println!("{:?}", res_2.await.status());
}

async fn request_login(client: Client, url: &str) -> Response {
    let res = client.post(url).send().await;

    match res {
        Ok(resp) => resp,
        Err(err) => {
            println!("{:?}", err);
            panic!("testing")
        }
    }
}
