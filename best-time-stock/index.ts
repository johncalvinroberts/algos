// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
function maxProfit(prices: number[]): number {
  let buyPrice = 0;
  let sellPrice = 1;
  let maxProfit = 0
  while (sellPrice < prices.length) {
    if (prices[sellPrice] > prices[buyPrice]) {
      maxProfit = Math.max(maxProfit, prices[sellPrice] - prices[buyPrice]);
    } else {
      buyPrice = sellPrice
    }
    sellPrice++;
  }
  return maxProfit;
};

(() => {
  console.log(maxProfit([7, 1, 5, 3, 6, 4]));// 5
  console.log(maxProfit([7, 6, 4, 3, 1]));// 0
  console.log(maxProfit([1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9]));// 9

})()