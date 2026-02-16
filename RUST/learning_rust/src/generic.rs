fn find_largest<T: std::cmp::PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];

    for item in list {
        if item > largest {
            largest = item;
        }
    }
    largest
}

pub fn generic_datatypes() {
    let array = vec![1, 45, 98, 90, 22];

    let largest = find_largest(&array);
    println!("the largest of array is : {largest}");

    let charact = vec!['c', 'b', 'z', 'o', 'y'];
    let largest = find_largest(&charact);
    println!("the largest of array is : {largest}");
}