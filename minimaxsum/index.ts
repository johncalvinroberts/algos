function miniMaxSum(arr: number[]): void {
  // Write your code here
  const sorted = arr.sort((a, b) => a - b);
  const front = sorted.slice(0, 4);
  const back = sorted.slice(sorted.length - 4, sorted.length);
  let sums = [0, 0];
  for (let i = 0; i < 4; i++) {
    sums[0] = sums[0] + front[i];
    sums[1] = sums[1] + back[i];
  }
  console.log(...sums);
}

const testCase = [1, 3, 5, 7, 9];
// expect print [16, 24]
console.log(miniMaxSum(testCase));
