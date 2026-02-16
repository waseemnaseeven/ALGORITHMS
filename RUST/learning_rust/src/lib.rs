pub mod struct_module;
pub mod collections_module;

pub fn only_variables() {
    println!("....Only variables....");
    let x: i32 = 8;
    println!("let x = {}", x);
    // x = 9;
    println!("cannot do x = 9 cuz its by default x is an immutable variable");
    let mut mutable_var: char = 'c';
    println!("mutable_var is {} and now we will change it to 'd'", mutable_var);
    mutable_var = 'd';
    println!("mutable_var is {}", mutable_var);
    let y: i32 = x;
    println!("let y = x so y = {}", y);
    let x = 12;
    println!("if i do let x = 12 so x = {}", x);
    const Z: i8 = 10;
    println!("const Z = {}", Z);
    let f: f64 = 3.64;
    println!("float is {}", f);

    println!("Lets deal with some groupments of types");
    let info = (1, 3.3, 2147483647, "coucou");
    let (jets, fuel, ammo, msg) = info;
    println!("{:?}", info);
    println!("there is all var in info = jets, fuel, ammo, msg {} {} {} {}", jets, fuel, ammo, msg);

}

pub fn volume(w: i32, h: i32, d: i32) -> i32 {
    println!("....Functions....");
    return w*h*d;
}

pub fn tuples_array() {
    let tuples: (u8, f64, f32, i8, i16, i32, i64) = (1, 3.33, 3.32, 8, 4, 45, 32);
    let array: [f64; 4] = [21479735.98888, 8.89678, 798.7386, 0.0];

    // Printing the tuple
    println!("\n....Tuples and Arrays....\n");
    println!("Tuple:");

    // Using positional indexing
    println!("Index 0: {}", tuples.0);
    println!("Index 1: {}", tuples.1);
    println!("Index 2: {}", tuples.2);
    println!("Index 3: {}", tuples.3);
    println!("Index 4: {}", tuples.4);
    println!("Index 5: {}", tuples.5);
    println!("Index 6: {}", tuples.6);

    // Using pattern matching for a loop
    println!("\nUsing a for loop with pattern matching:");
    let tuple_elements: [&dyn std::fmt::Display; 7] = [
        &tuples.0,
        &tuples.1,
        &tuples.2,
        &tuples.3,
        &tuples.4,
        &tuples.5,
        &tuples.6,
    ];

    for (i, value) in tuple_elements.iter().enumerate() {
        println!("Index {}: {}", i, value);
    }

    // Printing the array
    println!("\nArray:");

    // Using positional indexing
    println!("Using a positional indexing:");
    println!("Index 0: {}", array[0]);
    println!("Index 1: {}", array[1]);
    println!("Index 2: {}", array[2]);
    println!("Index 3: {}", array[3]);

    // Using a for loop with index
    println!("\nUsing a for loop with index:");
    for i in 0..array.len() {
        println!("Index {}: {}", i, array[i]);
    }

    // Using a for loop with iter
    println!("\nUsing a for loop with iter:");
    for (i, &value) in array.iter().enumerate() {
        println!("Index {}: {}", i, value);
    }
}

pub fn inspect(s: &String) { // prends une reference de string
    println!{"\n Here is the word, i sent a reference to my string and it gave me {}", s};
}

pub fn change(s: &mut String) {
    s.push_str("s");
    println!("I put an s at the end {}", s);
    let new_str = s.replace("Hello", "Buonjiorno");
    println!("{}", new_str);
    println!("I replace {}", s.replace("Hello", "Bonjour"));
}

pub fn ownership() {
    println!("\n....Ownership....");
    let mut mots = String::from("'Hello World'");
    inspect(&mots);
    change(&mut mots);
    let new_str = String::from("Coucou");
    let new_one = new_str;
    println!("new_one is : {} and cannot print the string from where it come from ", new_one);
    let new2 = new_one.clone();
    println!("i made a deep copy with .clone(), now its {} and new_one is available aswell : {} ", new2, new_one);
}

