function isValid(s: string): boolean {
  const stack: string[] = [];
  const opening = "([{";
  const closing = ")]}"
  const split = s.split("");
  while (split.length > 0) {
    const sub = split.shift();
    if (!sub) {
      break;
    }
    if (opening.includes(sub)) {
      stack.push(sub);
      continue;
    }
    if (closing.includes(sub)) {
      const o = stack.pop();
      if (!o) {
        return false
      }
      if (opening.indexOf(o) != closing.indexOf(sub)) {
        return false
      }
    }
  }
  if (stack.length != 0) {
    return false
  }
  return true
};

(() => {
  console.log(isValid("()"))// true
  console.log(isValid("()[]{}"))// true
  console.log(isValid("(]"))// false
  console.log(isValid("["))// false
})()