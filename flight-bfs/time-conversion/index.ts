/**
 * convert a 12-hr time sting to military time
 * e.g., 12:01:00AM -> 00:01:00
 *
 */

const timeRef = {
  AM: {
    "12": "00",
    "01": "01",
    "02": "02",
    "03": "03",
    "04": "04",
    "05": "05",
    "06": "06",
    "07": "07",
    "08": "08",
    "09": "09",
    "10": "10",
    "11": "11",
  },
  PM: {
    "12": "12",
    "01": "13",
    "02": "14",
    "03": "15",
    "04": "16",
    "05": "17",
    "06": "18",
    "07": "19",
    "08": "20",
    "09": "21",
    "10": "22",
    "11": "23",
  },
};

function timeConversion(s) {
  // Write your code here
  const amPm = s.substring(s.length - 2, s.length);
  const [hrs, ...rightSideOfTime] = s.substring(0, s.length - 2).split(":");
  console.log({ rightSideOfTime });
  return `${timeRef[amPm][hrs]}:${rightSideOfTime.join(":")}`;
}

const testCase = "07:05:45PM";
// expect 19:05:45
console.log(timeConversion(testCase));
