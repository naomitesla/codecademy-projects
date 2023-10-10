// Constant variable for the current temparature reading in kelvin c:
const kelvin = 293;

// Celsius is kelvin minus 273 degrees c:
let celsius = kelvin - 273;

// Formula for converting celsius to fahrenheit c:
let fahrenheit = Math.floor(celsius * (9/5) + 32);

let newton = Math.floor(celsius * (33/100));

console.log(`The temperature is ${fahrenheit} degrees Fahrenheit.`);
console.log(`The temperature is ${newton} degrees on the Newton scale.`);
