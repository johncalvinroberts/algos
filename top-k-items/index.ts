function topKFrequent(nums: number[], k: number): number[] {
  const counts = new Map<number, number>();
  while (nums.length > 0) {
    const item = nums.pop() ?? 0;
    counts.set(item, (counts.get(item) ?? 0) + 1);
  }
  return Array.from(counts).sort(([, a], [, b]) => b - a).slice(0, k).map(([key]) => key);
};

(() => {
  console.log(topKFrequent([1, 1, 1, 2, 2, 3], 2));// [1,2]
})()