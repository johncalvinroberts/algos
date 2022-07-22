function staircase(n: number): void {
  // Write your code here
  let stairs: string[] = [];
  for (let j = 0; j < n; j++) {
    const offset = n - j - 1;
    let step: string[] = [];
    for (let i = 0; i < n; i++) {
      if (i >= offset) step.push("#");
      else step.push(" ");
    }
    stairs.push(step.join(""));
  }
  console.log(stairs.join("\n"));
}

/**
expect this:

     #
    ##
   ###
  ####
 #####
######
 */
staircase(6);
