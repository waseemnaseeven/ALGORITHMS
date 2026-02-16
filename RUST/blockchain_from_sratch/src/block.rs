use std::time::SystemTime;
use crypto:: {
    digest::Digest,
    sha2::Sha256,
};
use log::info;
use serde::{Serialize, Deserialize};

use crate::blockchain::Blockchain;
use crate::errors::Result;
use crate::transaction::Transaction;

// Difficulty of the proof-of-work, its the number of 0 the hash has for a bloc to be valid
const TARGET_NEXT: usize = 4;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Block {
    timestamp: u128, // time the transaction was created
    transactions: Vec<Transaction>, 
    prev_block_hash: String,
    hash: String, // hash for himself
    height: usize, // an index 
    nonce: i32, // number of tries on the process 'proof-of-work'
}

impl Block {

    pub fn get_transaction(&self) -> &Vec<Transaction> {
        &self.transactions
    }

    pub(crate) fn get_prev_hash(&self) -> String {
        self.prev_block_hash.clone()
    }

    pub fn new_genesis_block(coinbase: Transaction) -> Block {
        Block::new_block(vec![coinbase], String::new(), 0).unwrap()
    }

    pub fn get_hash(&self) -> String {
        self.hash.clone()
    }

    pub fn new_block(data: Vec<Transaction>, prev_block_hash: String, height: usize) -> Result<Block> {
        let timestamp: u128 = SystemTime::now()
            .duration_since(SystemTime::UNIX_EPOCH)?
            .as_millis();
        
        let mut block = Block {
            timestamp: timestamp,
            transactions: data,
            prev_block_hash,
            hash: String::new(),
            height: height,
            nonce: 0,
        };

        block.run_proof_if_work()?;
        Ok(block)
    }

    fn run_proof_if_work(&mut self) -> Result<()> {
        info!("mining the block");
        while !self.validate()? {
            self.nonce += 1;
        }
        let data = self.prepare_hash_data()?;
        let mut hasher = Sha256::new();
        hasher.input(&data[..]);
        self.hash = hasher.result_str();
        Ok(())
    }

    fn validate(&self) -> Result<bool> {
        let data = self.prepare_hash_data()?;
        let mut hasher = Sha256::new();
        hasher.input(&data[..]);
        let hash_result = hasher.result_str();
        Ok(hash_result.starts_with(&"0".repeat(TARGET_NEXT)))
    }

    fn prepare_hash_data(&self) -> Result<Vec<u8>> {
        let content = (
            &self.prev_block_hash,
            &self.transactions,
            self.timestamp,
            TARGET_NEXT,
            self.nonce
        );
        let bytes = bincode::serialize(&content)?;
        Ok(bytes)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    // src/block.rs
    #[test]
    fn test_blockchain() {
        let mut my_fabulous_blockchain = Blockchain::new().unwrap();
    }
}
