// Copyright (c) 2023 Furukawa, Atsushi <atsushifx@aglabo.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT



function groupAnagrams(strs: string[]): string[][] {
  // functions
  const strHash = (function (str: string): string {
    return str.split("").sort().toString()
  });

  //
  const myHash: { [key: string]: string[] } = {}

  strs.map(str => {
    const h: string = strHash(str)
    if (!myHash.hasOwnProperty(h)) {
      myHash[h] = []
    }
    myHash[h].push(str)
  })
  return Object.values(myHash).sort()
};


// main roution

main()
function main() {
  let input: string[];
  let output: string[][];
  input = ["eat", "tea", "tan", "ate", "nat", "bat"]
  output = groupAnagrams(input)
  console.log(output);
}

