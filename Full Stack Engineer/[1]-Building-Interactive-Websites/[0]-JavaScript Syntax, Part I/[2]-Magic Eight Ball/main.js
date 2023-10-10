let username = 'naomi';
username ? console.log(`[*] hiii ${username}!!`) : console.log(`[*] umm hiii`);

console.log();

let userQuestion = 'will i ever find true happiness?';
console.log(` > ${username}:  ${userQuestion}`);

console.log();

let randomNumber = Math.floor(Math.random() * 8);
let eightBall = '';

switch (randomNumber){
  case 0:
    eightBall = 'umm for sureee c:';
    break;
  case 1:
    eightBall = 'nyoo im sorry :c';
    break;
  case 2:
    eightBall = 'probably not!';
    break;
  case 3:
    eightBall = 'i literally just hate you!!';
    break;
  case 4:
    eightBall = 'look elsewhere!';
    break;
  case 5:
    eightBall = 'idk :c';
    break;
  case 6:
    eightBall = 'probably c:';
    break;
  case 7:
    eightBall = 'get a life and stop talking to me';
    break;
  default:
    break;
}

eightBall ? console.log(eightBall) : console.log('how did you manage to mess this up?');

