import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: ms (beats % of TS solutions)
 *   Memory: mb (beats % of TS solutions)
 */

// bad solution i dont like it
function findMaxLength(nums: number[]): number {
    const cache: any = { 0: -1 };
    let max = 0;
    let count = 0;
    for (let i = 0; i < nums.length; i++) {
        nums[i] == 0 ? count-- : count++;
        if (count in cache) {
            max = Math.max(max, i - cache[count]);
        } else {
            cache[count] = i;
        }
    }
    return max;
}

// Test cases
assertEquals(2, findMaxLength([0, 1]));
assertEquals(2, findMaxLength([0, 1, 0]));
