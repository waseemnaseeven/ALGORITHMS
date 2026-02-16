const STARTING_MISSILES: i32 = 8;
const READY_AMOUNT: i32 = 2;

fn main() {
    println!("Hello, world!");

    let missiles = STARTING_MISSILES;
    let ready = READY_AMOUNT;
    READY_AMOUNT = 1;
    let finish = 0;
    println!("Firing {} of my {} missiles...", ready, missiles);

    println!("{} missiles left", missiles - ready);


}
