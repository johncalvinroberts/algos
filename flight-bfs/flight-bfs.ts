// https://leetcode.com/problems/cheapest-flights-within-k-stops/discuss/687049/Javascript-BFS

const findCheapestTrip = (
  flights: number[][],
  src: number,
  whereIWantToGo: number
) => {
  // find cheapest flight or combination of flights
  const graph: Record<number, Record<number, number>> = {};
  for (let [source, destination, price] of flights) {
    if (!graph[source]) {
      graph[source] = {};
    }
    graph[source][destination] = price;
  }
  let queue = [[0, src]];
  while (queue.length) {
    let [price, destination] = queue.shift() || [];
    if (destination === whereIWantToGo) return price;
    const nextStops = graph[destination];
    if (nextStops) {
      for (let connectionString in nextStops) {
        const connection = parseInt(connectionString);
        let cost = graph[destination][connection];
        queue.push([cost + price, connection]);
      }
    }
    queue.sort((a, b) => a[0] - b[0]);
  }
  return -1;
};

const testFlights = [
  [0, 1, 100],
  [1, 2, 100],
  [2, 3, 100],
  [0, 3, 500],
  [1, 3, 600],
];
const takeoff = 0;
const destination = 3;

// expect 300
console.log(findCheapestTrip(testFlights, takeoff, destination));
