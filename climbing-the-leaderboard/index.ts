function climbingLeaderboard(
  ranked: number[],
  playerScores: number[]
): number[] {
  const rankedSet = Array.from(new Set(ranked));
  const rankMap = new Map<number, number>();
  let ranks: number[] = [];
  for (const s of playerScores) {
    let rank = 1;
    const cached = rankMap.get(s);
    if (cached) {
      rank = cached;
      continue;
    }
    for (let i = rankedSet.length; i > 0; i--) {
      if (s < rankedSet[i - 1]) {
        rank = i + 1;
        break;
      }
    }
    ranks.push(rank);
  }
  return ranks;
}

const testRanked = [100, 100, 50, 40, 40, 20, 10];

const testPlayerScores = [5, 25, 50, 120];
// expect [6,4,2,1]
console.log(climbingLeaderboard(testRanked, testPlayerScores));
