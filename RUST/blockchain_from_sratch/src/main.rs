mod block;
mod blockchain;
mod errors;
mod cli;

use crate::cli::Cli;

pub type Result<T> = std::result::Result<T, failure::Error>;

fn main() -> Result<()> {
    println!("Blockchain fromn Scratch is running....");
    let mut cli = Cli::new()?;
    cli.run()?;
    Ok(())
}