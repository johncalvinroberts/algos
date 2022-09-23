function calculate(s) {
  let operator = "+";
  let currentNum = 0;
  let numStack = [];

  for (let i = 0; i <= s.length; i++) {
    let char = s[i];
    // handle space case
    if (char === " ") continue;
    console.log({ numStack, currentNum, operator, v: char });
    // if it's number, add to currentNum
    if (!isNaN(char)) {
      currentNum = currentNum * 10 + parseInt(char);
      // if it's operator, make calculation and add to numStack
    } else {
      if (operator === "+") {
        numStack.push(currentNum);
      } else if (operator === "-") {
        numStack.push(-currentNum);
      } else if (operator === "*") {
        let num = numStack.pop();
        console.log({ num, currentNum });
        numStack.push(num * currentNum);
      } else if (operator === "/") {
        let num = numStack.pop();
        numStack.push(Math.trunc(num / currentNum));
      }
      // reset num
      currentNum = 0;
      // assing new opertor
      operator = char;
    }
  }
  return numStack.reduce((a, b) => a + b);
}

(() => {
  const cases = [
    // ["3+2*2", 7],
    // [" 3/2 ", 1],
    // [" 3+5 / 2 ", 5],
    ["3+3*10-2", 31],
    // ["4*10+2", 42],
  ];
  for (const [key, value] of cases) {
    const result = calculate(key);
    if (result !== value) {
      console.log(`FAILED: ${key}, got ${result} expected ${value}`);
    } else {
      console.log(`PASSED: ${key}, got ${result} expected ${value}`);
    }
  }
})();
