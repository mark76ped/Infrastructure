// use rand::Rng;
// use rand::thread_rng;
use rust_project::on_off;
use rust_project::print_difference;
use rust_project::ding;
// const STARTING_MISSILES: i32 = 8;
// const READY_AMOUNT: i32 = 2;

fn main() {
//     let mut missiles = STARTING_MISSILES;
//     let ready = READY_AMOUNT;
//     println!("Firing {} of my {} missiles...", ready, missiles);
//     missiles = missiles - ready;
//     println!("{} missiles left", missiles);

//     let width = 4;
//     let height = 7;
//     let depth = 10;
//     // 1. Try running this code with `cargo run` and take a look at the error.
//     //
//     // See if you can fix the error. It is right around here, somewhere.  If you succeed, then
//     // doing `cargo run` should succeed and print something out.
    
//     let area = area_of(width, height);
    
//     println!("Area is {}", area);

//     // 2. The area that was calculated is not correct! Go fix the area_of() function below, then run
//     //    the code again and make sure it worked (you should get an area of 28).

//     // 3. Uncomment the line below.  It doesn't work yet because the `volume` function doesn't exist.
//     println!("Volume is {}", volume(width, height, depth));

//     let seven = thread_rng().gen_range(0,100);
//     println!("The random number between 1 and 100 is {}", seven);
    let coords: (f32, f32) = (6.3, 15.0);
    print_difference(coords.0, coords.1);

    // let coords_arr: [f32; 2] = [coords.0, coords.1];            // create an array literal out of parts of `coord` here
    // print_array(coords_arr);        // and pass it in here (this line doesn't need to change)
    
    
    let series = [1, 1, 2, 3, 5, 8, 13];
    ding(series[6]);

    let mess = ([3, 2], 3.14, [(false, -3), (true, -100)], 5, "candy");
     //println!("{}", mess.2[1].0);
     on_off(mess.2[1].0);
}
