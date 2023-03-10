function isAnagram(s: string, t: string): boolean {
  if (s.length != t.length) return false;
  const arr = s.split("");
  while (arr.length > 0) {
    const l = arr.shift();
    if (!l) {
      continue;
    }
    if (!t.includes(l)) {
      return false;
    } else {
      t = t.replace(l, "");
    }
  }
  return true;
}

(() => {
  console.log(isAnagram("anagram", "nagaram"));
  console.log(isAnagram("rat", "cat"));
  console.log(isAnagram("cacc", "caca"));
})();
