use crate::block::Block;
use crate::errors::Result;

use bincode::{deserialize, serialize};
use std::collections::HashMap;

use log::info;

const TARGET_NEXT: usize = 4;
const GENESIS_COINBASE_DATA: &str =
    "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks";

#[derive(Debug, Clone)]
pub struct Blockchain {
    current_hash: String,
    db: sled::Db, // stores the blocks, with a key-value librairy
}

// Iterate over the blockchain thanks to current_hash
pub struct BlockchainIter<'a> {
    current_hash: String, // hash of the actual bloc
    bc: &'a Blockchain,
}

impl Blockchain {
    // Create a new Blockchain, database creation 
    pub fn new() -> Result<Blockchain> {
        info!("Open a blockchain");

        let db = sled::open("data/blocks")?;
        let hash = db
            .get("LAST")?
            .expect("Must create a new block database first");
        info!("Found block database");
        let lasthash = String::from_utf8(hash.to_vec())?;
        Ok(Blockchain {
            current_hash: lasthash.clone(),
            db,
        })
    }

    pub fn create_blockchain(address: String) -> Result<Blockchain> {
        info!("Creating a new blockchain");

        let db = sled::open("data/blocks")?;
        info!("Creating new block database");
        let cbtx =  Transaction::new_coinbase(address, String::from(GENESIS_COINBASE_DATA))?;
        let genesis: Block = Block::new_genesis_block(cbtx);
        db.insert(genesis.get_hash(), serialize(&genesis)?)?;
        db.insert("LAST", genesis.get_hash().as_bytes())?;
        let bc = Blockchain {
            current_hash: genesis.get_hash(),
            db,
        };
        bc.db.flush()?;
        Ok(bc)
    }

    pub fn add_block(&mut self, transactions: Vec<Transaction>) -> Result<()> {
        let last_hash = self.db.get("LAST")?.unwrap(); // Get the last hash

        let new_block = Block::new_block(transactions, String::from_utf8(last_hash.to_vec())?, TARGET_NEXT)?; // create new block with lasthash and the difficulty
        self.db.insert(new_block.get_hash(), bincode::serialize(&new_block)?)?; // adding into the db
        self.db.insert("LAST", new_block.get_hash().as_bytes())?; // put into 'LAST' the last hash, meaning the hash of the actual block we just created
        self.current_hash = new_block.get_hash(); // updating current_hash to point into the last hash
        Ok(())
    }

    // FindUnspentTransactions returns a list of transactions containing unspent outputs
    fn find_unspent_transactions(&self, address: &str) -> Vec<Transaction> {
        let mut spent_TXOs: HashMap<String, Vec<i32>> = HashMap::new();
        let mut unspend_TXs: Vec<Transaction> = Vec::new();

        for block in self.iter() {
            for tx in block.get_transaction() {
                for index in 0..tx.vout.len() {
                    if let Some(ids) = spent_TXOs.get(&tx.id) {
                        if ids.contains(&(index as i32)) {
                            continue;
                        }
                    }

                    if tx.vout[index].can_be_unlock_with(address) {
                        unspend_TXs.push(tx.to_owned())
                    }
                }

                if !tx.is_coinbase() {
                    for i in &tx.vin {
                        match spend_txos.get_mut(&i.txid) {
                            Some(v) => {
                                v.push(i.vout);
                            }
                            None => {
                                spend_txos.insert(i.txid.clone(), vec![i.vout]);
                            }
                        }
                    }
                }
            }
        }
        unspend_TXs
    }

    pub fn find_UTXO(&self, address: &str) -> Vec<TXOutput> {
        let mut utxos = Vec::<TXOutput>::new();
        let unspend_TXs = self.find_unspent_transactions(address);
        for tx in unspend_TXs {
            for out in &tx.vout {
                if out.can_be_unlock_with(&address) {
                    utxos.push(out.clone());
                }
            }
        }
        utxos
    }

    pub fn find_spendable_outputs (
        &self,
        address: &str,
        amount: i32,
    ) -> (i32, HashMap::<String, Vec<i32>>) {
        let mut unspent_outputs = HashMap::<String, Vec<i32>> = HashMap::new();
        let mut accumulated = 0;
        let unspend_TXs = self.find_unspent_transactions(address);

        for tx in unspend_TXs {
            for index in 0..tx.vout.len() {
                if tx.vout[index].can_be_unlock_with(address) && accumulated < amount {
                    match unspent_outputs.get_mut(&tx.id) {
                        Some(v) => v.push(index as i32),
                        None => {
                            unspent_outputs.insert(tx.id.clone(), vec![index as i32]);
                        }
                    }
                    accumulated += tx.vout[index].value;

                    if accumulated >= amount {
                        return (accumulated, unspent_outputs);
                    }
                }
            }
        }
        (accumulated, unspent_outputs)
    }

    // Send an iterator of the last hash
    pub fn iter(&self) -> BlockchainIter {
        BlockchainIter {
            current_hash: self.current_hash.clone(),
            bc: &self,
        }
    }
}

// Implement how the iterator should behave, meaning block by block
impl<'a> Iterator for BlockchainIter<'a> {
    type Item = Block;

    fn next(&mut self) -> Option<Self::Item> {
        if let Ok(encode_block) = self.bc.db.get(&self.current_hash) { // get the last hash 
            return match encode_block {
                Some(b) => { // if we find it, we return it, deserialize
                    if let Ok(block) = bincode::deserialize::<Block>(&b) {
                        self.current_hash = block.get_prev_hash();
                        Some(block) 
                    } else {
                        None
                    }
                }
                None => None,
            }
        } 
        None
    }
}

#[cfg(test)]
mod tests {
    // src/blockchain.rs
    use super::*;

    #[test]
    fn test_add_block() {
        let mut my_fabulous_blockchain = Blockchain::new().unwrap();

        for item in my_fabulous_blockchain.iter() {
            println!("item {:?}", item);
        }
    }
}