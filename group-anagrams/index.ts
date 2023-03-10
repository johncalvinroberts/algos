function groupAnagrams(strs: string[]): string[][] {
  const angrms: Map<string, string[]> = new Map();
  for (const str of strs) {
    const key = str
      .split("")
      .sort((a, b) => a.charCodeAt(0) - b.charCodeAt(0))
      .join("");
    angrms.set(key, [...(angrms.get(key) || []), str]);
  }
  return Array.from(angrms.values());
}

(() => {
  console.log(groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])); // [["bat"],["nat","tan"],["ate","eat","tea"]]
})();
