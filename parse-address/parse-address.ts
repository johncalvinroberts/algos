/**
 *     
 { address_lines_1: '60 E 54th St', postal_code: 10022, city: 'New York', country: 'USA' }
  { address_lines_1: '7 Agirre Miramon Kalea', postal_code: 20003, city: 'Donostia', country: 'Spain' }
  { address_lines_1: 'Refshalevej 96', postal_code: 1432, city: 'København K', country: 'Denmark' }
 */

const cases = [
  "60 E 54th St, New York, NY 10022, USA",
  "Agirre Miramon Kalea, 7, 20003 Donostia, Gipuzkoa, Spain",
  "Refshalevej 96, 1432 København K, Denmark",
];

const NUMBER_PLUS_WHITESPACE = /\d+ | \d+/;
const POSTAL_CODE_REGEX = /\d+/;

function parseAddress(
  rawAddress: string
): Record<string, string | number | undefined> {
  let address_lines_1: string = "";
  let postal_code: number = 0;
  let city: string = "";
  let country: string = "";
  const split = rawAddress.split(", ");
  let initialLength = split.length;
  while (split.length > 0) {
    const chunk = split.shift();
    if (split.length == initialLength - 1 && chunk) {
      let maybeStreetNumber = chunk.match(NUMBER_PLUS_WHITESPACE)?.[0];
      const street = chunk.replace(maybeStreetNumber || "", "");
      if (maybeStreetNumber == null) {
        maybeStreetNumber = split.shift();
      }
      address_lines_1 = `${maybeStreetNumber} ${street}`;
    }
  }

  return {
    address_lines_1,
    postal_code,
    city,
    country,
  };
}

console.log(parseAddress(cases[0]));
console.log(parseAddress(cases[1]));
console.log(parseAddress(cases[2]));
