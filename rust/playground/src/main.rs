struct Person {
    age: u8,
    name: String,
}

impl Person {
    fn new(age: u8, name: String) -> Person {
        Person {
            age: age,
            name: name,
        }
    }
    fn print_age(&self) {
        println!("{}", self.age);
    }
}

fn main() {
    let jin = Person {
        age: 16,
        name: String::from("Seungjin"),
    };

    println!("Print out person {} {}", jin.age, jin.name);
}
