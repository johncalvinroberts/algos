function isPalindrome(s: string): boolean {
  const massaged = s.replace(/[^\w+]|[_]/g, "").toLowerCase();
  console.log({ massaged })
  for (let i = 0; i < massaged.length; i++) {
    if (massaged[i] != massaged[massaged.length - 1 - i]) {
      return false
    }
  }
  return true
};

(() => {
  console.log(isPalindrome("racecar"));//true
  console.log(isPalindrome("A man, a plan, a canal: Panama"));//true
  console.log(isPalindrome("kitty"));//false
  console.log(isPalindrome("ab_a"));//true
})()