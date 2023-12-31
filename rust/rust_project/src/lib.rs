pub fn print_difference(x: f32, y: f32) {
    println!("Difference between {} and {} is {}", x, y, (x - y).abs());
}

pub fn on_off(val: bool) {
    if val {
        println!("Lights are on!");
    }
}

pub fn ding(x: i32) {
    if x == 13 {
        println!("Ding, you found 13!");
    }
}

pub fn print_array(a: [f32; 2]) {
    println!("The coordinates are ({}, {})", a[0], a[1]);
}


pub fn area_of(x: i32, y: i32) -> i32 {
    x * y
}

pub fn volume(x: i32, y: i32, z: i32) -> i32 {
    x * y * z
}
