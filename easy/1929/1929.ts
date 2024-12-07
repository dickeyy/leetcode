import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 68ms (beats 67.06% of TS solutions)
 *   Memory: 54.23mb (beats 41.42% of TS solutions)
 */
function getConcatenation(nums: number[]): number[] {
    return nums.concat(nums);
}

// test cases
assertEquals([1, 2, 1, 1, 2, 1], getConcatenation([1, 2, 1]));
assertEquals([1, 3, 2, 1, 1, 3, 2, 1], getConcatenation([1, 3, 2, 1]));
