use std::fs;

fn main() {
    let input = fs::read_to_string("day01/input.txt").unwrap();

    let part1 = input.chars().fold(0, |acc, c| match c {
        '(' => acc + 1,
        _ => acc - 1,
    });

    let mut floor = 0;
    let mut part2 = 0;

    for (i, c) in input.chars().enumerate() {
        match c {
            '(' => floor += 1,
            _ => floor -= 1,
        };
        if floor == -1 {
            part2 = i;
            break;
        }
    }

    println!("part1: {part1}");
    println!("part2: {part2}");
}
