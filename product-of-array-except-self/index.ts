function productExceptSelf(nums: number[]): number[] {
  const ret: number[] = [];
  for (let i = 0, product = 1; i < nums.length; i++) {
    ret[i] = product;
    product = product * nums[i];
  }
  for (let i = nums.length - 1, product = 1; i >= 0; i--) {
    ret[i] *= product;
    product *= nums[i];
  }
  return ret;
};


(() => {
  console.log(productExceptSelf([1, 2, 3, 4]));//[24,12,8,6]
  console.log(productExceptSelf([-1, 1, 0, -3, 3]));//[0,0,9,0,0]
})();