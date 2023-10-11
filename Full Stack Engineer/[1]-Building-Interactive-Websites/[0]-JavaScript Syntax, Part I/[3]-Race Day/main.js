let raceNumber = Math.floor(Math.random() * 1000);
let early = false;
let age = 22;

let adult = age > 18;

if (early && adult) {
    raceNumber += 1000;
    console.log(`[${raceNumber}] ur racee starts at 9:30am!`);
} else if (!early && adult) {
    console.log(`[${raceNumber}] ur racee starts at 11:00am!`);
} else if (age == 18) {
    console.log('umm ig u have to go see the registration desk :c')
} else {
    console.log(`[${raceNumber}] ur racee starts at 12:30am!`);
}
