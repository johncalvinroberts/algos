function twoSum(nums: number[], target: number): number[] {
  const numsEnc: (number | null)[] = [...nums];
  for (let i = 0; i < nums.length; i++) {
    const n = nums[i];
    numsEnc[i] = null;
    const numberNeeded = target - n;
    const found = numsEnc.indexOf(numberNeeded);
    if (found != -1) {
      return [i, found];
    }
  }
  return [];
}

(() => {
  console.log(twoSum([2, 7, 11, 15], 9)); // [0,1]
  console.log(twoSum([3, 2, 4], 6)); // [1,2]
  console.log(twoSum([3, 3], 6)); // [0,1]
})();
