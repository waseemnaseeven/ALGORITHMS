pub mod struct_module;

use crate::struct_module::everything_about_structs_traits;

use learning_rust::collections_module::some_collections;

use crate::generic::generic_datatypes;

use learning_rust::{only_variables, volume, tuples_array, ownership};

mod generic;

fn main() {
    // Variables
    only_variables();
    println!();
    // Parameters on functions
    println!("I called volume functions and gave 3 parameters i32 {}", volume(32,64,88));
    // Tuples and array 
    tuples_array();
    // Ownership
    ownership();
    // Struct and Traits
    everything_about_structs_traits();
    // Collections
    some_collections();
    // Generic Data Types aka Templates
    generic_datatypes();
}
