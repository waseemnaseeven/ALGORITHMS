use crate::blockchain::Blockchain;
use crate::errors::Result;
use crate::transaction::Transaction;

use clap::{
    arg,
    Command,
};

pub struct Cli {
    bc: Blockchain,
}

impl Cli {
    // Create a new blockchain or load the existing one
    pub fn new() -> Result<Cli> {
        let bc = Blockchain::new()?;
        Ok(Cli { bc })
    }

    pub fn run(&mut self) -> Result<()> {
        let matches = Command::new("blockchain-rust-demo")
            .version("0.1")
            .author("@gmail.com")
            .about("blockchain in rust: a simple blockchain for learning")
            .subcommand(Command::new("printchain").about("print all the chain blocks"))
            .subcommand(Command::new("addblock").about("add a new block to the chain").arg(arg!(<DATA>"'the blockchain data'")))
            .get_matches();
        if let Some(ref matches) = matches.subcommand_matches("addblock") {
            if let Some(c) = matches.get_one::<String>("DATA") {
                self.addblock(String::from(c))?;
            } else {
                println!("Not printing testing lists...");
            }
        }
        if let Some(_) = matches.subcommand_matches("printchain") {
            self.print_chain();
        }
        Ok(())
    }

    fn addblock(&mut self, data: String) -> Result<()> {
        self.bc.add_block(vec![])
    }

    fn print_chain(&mut self) {
        for b in &mut self.bc.iter() {
            println!("block: {:#?}", b);
        }
    }
}