use clap::{ Parser};


#[derive(Parser,Debug)]
#[command(version, about, long_about = None)]
struct Args {
    #[arg(short , long )]
    name : String,
    username: String,

    #[arg(short, long, default_value_t=1)]
    count: u8,
}

fn main() {

    let args = Args::parse();

    
    for _ in 0..args.count{
        println!("Hello {}!", args.name);
    }

    match args.name.as_str(){
        "token"=>
        {
            if args.count < 1 {
                    println!("Args is wrong please check again");
                    println!("args count is {}",args.count);
                    return;
                }
                let username = args.username;
                token_function(username)
        }
        
        _=> println!("nothing injected"),
    }

    
}

fn token_function(username: String){

    println!("Hello {}",username.as_str());

}