import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 50ms (beats 97.01% of TS solutions)
 *   Memory: 51.84mb (beats 83.37% of TS solutions)
 */

function missingNumber(nums: number[]): number {
    const n = nums.length;
    const sum = nums.reduce((acc, curr) => acc + curr, 0);
    const res = (n * n + n) / 2; // (n(n+1)/2)
    return res - sum;
}

// Test cases
assertEquals(2, missingNumber([3, 0, 1]));
assertEquals(2, missingNumber([0, 1]));
assertEquals(8, missingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1]));
assertEquals(0, missingNumber([1]));
