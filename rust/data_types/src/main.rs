fn main() {
    let x: i8 = 10;
    println!("{}", x);

    let y: u8 = 10;
    println!("{}", y);
 
    let decimal = 02_55;
    let hex = 0xff;
    let octal = 0o377;
    let binary = 0b1111_1111;

    println!("{}",decimal);
    println!("{}",hex);
    println!("{}",octal);
    println!("{}",binary);

    let byte = b'A';
    println!("{}",byte);

    let _x = 2.0;
    let _y: f32 = 1.0;


    let _t = true;
    let _f: bool = false;


    let c = 'c';
    println!("{}",c);

    let a = 11;
    let b = 2;

    let remainder = a % b;
    println!("{}", remainder);

    let tup = (500, "hi", true);
    println!("{}", tup.1);
    let (x, y, z) = tup;

    println!("{}", x);
    println!("{}", y);
    println!("{}", z);

    let array = [1,2,3];
    println!("{}", array[0]);  

    let mut array2: [i32; 3] = [4,5,6];
    println!("{}", array2[2]);

    array2[0] = 10;
    println!("{:?}", array2);

    let mut nums = vec![1,2,3];

    nums.push(4);
    println!("{:?}", nums);
    nums.pop();
    println!("{:?}", nums);

    let mut vec = Vec::new(); //vec!
    vec.push("Test");
    vec.push("String");
    println!("{:?}", vec);

    vec.reverse();
    println!("{:?}", vec);

    let vect = Vec::<i32>::with_capacity(2);
    println!("{}", vect.capacity());

    let v : Vec<i32> = (0..5).collect();
    println!("{:?}", v);

    let sv: &[i32] = &v[2..4];
    println!("{:?}", sv); 

    let name = String::from("Mark");
    let course = "Rust".to_string();
    let new_name = name.replace("Mark", "Chad");

    println!("{:?}", name);
    println!("{:?}", course);
    println!("{:?}", new_name);

    let str1 = "hello";
    println!("{}", str1);
    let str2 = str1.to_string();
    let str3 = &str2;

    println!("{:?}", str1);
    println!("{:?}", str2);
    println!("{:?}", str3);
    println!("{}", "ONE".to_lowercase() == "one");



}
