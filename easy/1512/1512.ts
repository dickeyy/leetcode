import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 42ms (beats 98.44% of TS solutions)
 *   Memory: 50.26mb (beats 43.33% of TS solutions)
 */

function numIdenticalPairs(nums: number[]): number {
    let count = 0;
    for (let i = 0; i < nums.length; i++) {
        for (let j = 0; j < nums.length; j++) {
            if (nums[i] == nums[j] && i < j) count++;
        }
    }
    return count;
}

// Test cases
assertEquals(4, numIdenticalPairs([1, 2, 3, 1, 1, 3]));
assertEquals(6, numIdenticalPairs([1, 1, 1, 1]));
assertEquals(0, numIdenticalPairs([1, 2, 3]));
