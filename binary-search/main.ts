// https://leetcode.com/problems/binary-search/

function search(nums: number[], target: number): number {
  let l = 0;
  let r = nums.length;
  while (l <= r) {
    const middle = Math.round((l + r) / 2);
    if (nums[middle] > target) {
      r = middle - 1;
    }
    if (nums[middle] < target) {
      l = middle + 1;
    }
    if (nums[middle] == target) {
      return middle
    }
  }
  return -1
};

(() => {
  // console.log(search([-1, 0, 3, 5, 9, 12], 9))// 4, index
  // console.log(search([-1, 0, 3, 5, 9, 12], 2))// -1, index
  console.log(search([5], 5))// 0, index
})()