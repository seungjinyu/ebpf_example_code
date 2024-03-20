use select::document::Document;

use libc;

use std::{env, io, thread};

use reqwest::{self};

fn main() {
    let url = "https://www.google.com/search?q=computer+science+paper";

    let args: Vec<String> = env::args().collect();

    // let num_cores = num_cpus::get();

    // println!("number of cores {:?}", num_cores);

    // let handle = thread::spawn(move || {
    //     let cord_id = affinity::get_thread_affinity();
    //     println!("Hello from CPU core {}", cord_id.unwrap());
    // });
    let other_thread = thread::spawn(|| {
        println!(
            "I am thread {:?} on cpu {:?}",
            thread::current().id(),
            current_cpu()
        );
    });
    println!(
        "I am thread {:?} on cpu {:?}",
        thread::current().id(),
        current_cpu()
    );
    other_thread.join().unwrap();

    match args.get(1).unwrap().trim() {
        "single" => {
            let result = request_page(url);
            match result {
                Ok(res) => println!("{}", res),
                Err(err) => println!("{}", err),
            }
        }
        "web" => {
            let result = request_page_and_sort(url);
            match result {
                Ok(titles) => {
                    for title in titles {
                        println!("{}", title);
                    }
                }
                Err(err) => println!("Error {}", err),
            }
        }
        _ => {
            println!("No options")
        }
    };
}

// fn request_page(url: &str) -> Result<Vec<String>, reqwest::Error> {
fn request_page(url: &str) -> Result<String, reqwest::Error> {
    let resp = reqwest::blocking::get(url)?;
    let body = resp.text()?;

    println!("Request successful");

    Ok(body)
}
fn request_page_and_sort(url: &str) -> Result<Vec<String>, reqwest::Error> {
    let resp = reqwest::blocking::get(url)?;
    let body = resp.text()?;
    let titles = parse_titles(&body);

    println!("Request successful");
    Ok(titles)
}

fn parse_titles(html: &str) -> Vec<String> {
    let mut titles = Vec::new();

    let document = Document::from_read(html.as_bytes()).unwrap();

    for node in document.find(select::predicate::Name("dURPMd")) {
        println!("{}", node.text());
        let title = node.text();
        titles.push(title.trim().to_string());
    }

    titles
}

fn current_cpu() -> Result<usize, io::Error> {
    let ret = unsafe {
        libc::sched_getcpu();
    };

    if ret < 0 {
        Err(io::Error::last_os_error())
    } else {
        Ok(ret as usize)
    }
}
