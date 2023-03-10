function validateChunk(chunk: string[]): boolean {
  const items = new Map<number, number>();
  for (const item of chunk) {
    if (/[0-9]/.test(item)) {
      const count = items.get(parseInt(item)) || 0;
      items.set(parseInt(item), count + 1)
    }
  }
  if (Array.from(items.values()).some(item => item > 1)) {
    return false
  }
  return true
}

function isValidSudoku(board: string[][]): boolean {
  for (const row of board) {
    if (!validateChunk(row)) return false;
  }
  for (let i = 0; i < board.length; i += 3) {
    const chunk: string[] = []
    for (let j = i; j < 3; j++) {
      chunk.push(board[i][j]);
    }
    validateChunk(chunk)
  }
  return true
};

(() => {
  console.log(isValidSudoku([["5", "3", ".", ".", "7", ".", ".", ".", "."]
    , ["6", ".", ".", "1", "9", "5", ".", ".", "."]
    , [".", "9", "8", ".", ".", ".", ".", "6", "."]
    , ["8", ".", ".", ".", "6", ".", ".", ".", "3"]
    , ["4", ".", ".", "8", ".", "3", ".", ".", "1"]
    , ["7", ".", ".", ".", "2", ".", ".", ".", "6"]
    , [".", "6", ".", ".", ".", ".", "2", "8", "."]
    , [".", ".", ".", "4", "1", "9", ".", ".", "5"]
    , [".", ".", ".", ".", "8", ".", ".", "7", "9"]]));// true
  console.log(isValidSudoku([["8", "3", ".", ".", "7", ".", ".", ".", "."]
    , ["6", ".", ".", "1", "9", "5", ".", ".", "."]
    , [".", "9", "8", ".", ".", ".", ".", "6", "."]
    , ["8", ".", ".", ".", "6", ".", ".", ".", "3"]
    , ["4", ".", ".", "8", ".", "3", ".", ".", "1"]
    , ["7", ".", ".", ".", "2", ".", ".", ".", "6"]
    , [".", "6", ".", ".", ".", ".", "2", "8", "."]
    , [".", ".", ".", "4", "1", "9", ".", ".", "5"]
    , [".", ".", ".", ".", "8", ".", ".", "7", "9"]]));// false
})()