use std::collections::HashMap;


pub fn some_collections() {

    println!("\n----- Vectors -----\n");

    let mut v: Vec<i32> = Vec::new();
    println!("{:?}", v);
    v.push(4);
    println!("{:?}", v);

    let vect = vec![1,2,3,4,5];
    let third: Option<&i32> = vect.get(2);
    match third {
        Some(third) => println!("the third element is {third}"),
        None => println!("there is no third element"),
    }

    let v = vec![800,77,3];

    for num in 0..v.len() {
        println!("Index : {} = {} ", num, v[num]);
    }

    let u = vec![100, 32, 57];
    for i in &u {
        println!("{i}");
    }

    let mut adding = vec![100, 32, 57];
    for i in &mut adding {
        *i += 50;
        println!("{i}");
    }

    println!("\n\n----- Strings -----\n");

    let s1 = String::from("Hello, ");
    let s2 = String::from("world!");
    let s3 = s1 + &s2; // note s1 has been moved here and can no longer be used
    println!("We did s1 + s2 = s3: {} ", s3);

    let c1 = String::from("tic");
    let c2 = String::from("tac");
    let c3 = String::from("toe");

    let c = format!("{c1}-{c2}-{c3}");
    println!("{} ", c);

    let slice = &s3[1..10];
    println!("{slice}");
    let len: usize = s3.len();
    println!("{len}");

    println!("\n\n----- HashMap -----\n");

    let mut scores = HashMap::new();

    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    for (key, value) in &scores {
        println!("{key}: {value}");
    }

    let text = "hello world wonderful world";

    let mut map = HashMap::new();

    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }

    println!("{map:?}\n\n");

}