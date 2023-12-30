fn main() {
    
//     println!("{}", gcd(20, 15));
//     let one = 5;
//     if one > 4  {
//         println!("True");
//     } else if one == 1 {
//         println!{"equal"};
//     } else {
//         println!("False");
//     }
//     let mut num = 0;
//     'counter: loop {
//         println!("Count: {}", num);
//         let mut decrease = 5;
//         loop {
//             println!("Decreasing: {}", decrease);
//             if decrease == 4 {
//                 break;
//             }
//             if num == 2 {
//                 break 'counter;
//             }
//             decrease -= 1; 
//         }
//         num += 1;
//     }

//     let mut num = 0;
//     while num < 5 {
//         println!("Num: {}", num);
//         num += 1;
//     }
//         let vec: Vec<i8> = (0..10).collect();

//         for element in vec {
//             println!("{}", element); 
//         }
//         let val1 = 5;

//         let val2 = 2;

//         let ans = val1 % val2;
//         println!("{}", ans);
    let mut nums = vec![2, 4, 6, 8, 10];
    println!("{:?}", nums);
    nums.pop();
    nums.push(12);
    println!("{:?}", nums);

    let str1 = "Hello";
    concat_string(str1);

    let num1 = 26;
    control_flow(num1); 


}

fn concat_string(phrase: &str){
    println!("{} World", phrase);
}

fn control_flow (num: i32){
    if num == 1 {
        println!("The value is one"); 
    } else if num > 50 {
        println!("The value is greater than 50"); 
    } else if num < 25 {
        println!("The value is less than 25"); 
    } else {
        println!("The value is greater than 25 but less than 50"); 
    }
}

    // fn print_phase(phrase: &str){
//     println!("{}", phrase);
// }

//     fn gcd(mut a: u64, mut b: u64) -> u64 {
//         while a != 0{
//             if a < b {
//                 let c = a;
//                 a = b;
//                 b = c;
//             }
//             a = a % b;
//         }
//         b
//     }