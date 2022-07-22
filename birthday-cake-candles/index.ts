function birthdayCakeCandles(candles: number[]): number {
  let count = 0;
  let max = 0;
  for (let i = 0; i < candles.length; i++) {
    if (candles[i] > max) {
      max = candles[i];
      count = 0;
    }
    if (candles[i] === max) {
      count++;
    }
  }
  return count;
}

const testCase = [3, 2, 1, 3];
// expect 2
console.log(birthdayCakeCandles(testCase));
