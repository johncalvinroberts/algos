function isValidSudoku(board: string[][]): boolean {
  const rows = new Map<string, string>();
  const cols = new Map<string, string>();
  const squares = new Map<string, string>();
  for (let r = 0; r < 9; r++) {
    for (let c = 0; c < 9; c++) {
      const cell = board[r][c];
      if (cell === ".") continue;
      const grid = `${Math.floor(r / 3)} ${Math.floor(c / 3)}`;
      if (!rows[r]) rows[r] = new Set();
      if (!cols[c]) cols[c] = new Set();
      if (!squares[grid]) squares[grid] = new Set();
      if (rows[r].has(cell)) return false;
      if (cols[c].has(cell)) return false;
      if (squares[grid].has(cell)) return false;
      rows[r].add(cell);
      cols[c].add(cell);
      squares[grid].add(cell);
    }
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