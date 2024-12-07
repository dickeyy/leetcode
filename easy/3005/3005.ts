import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 65ms (beats 65.90% of TS solutions)
 *   Memory: 52.22mb (beats 35.84% of TS solutions)
 */

function maxFrequencyElements(nums: number[]): number {
    const freq = new Map();
    let max = 0;
    for (let i = 0; i < nums.length; i++) {
        const g = freq.get(nums[i]) + 1 || 1;
        if (max < g) max = g;
        freq.set(nums[i], g);
    }
    return Array.from(freq).filter((e) => e[1] === max).length * max;
}

// test cases
assertEquals(4, maxFrequencyElements([1, 2, 2, 3, 1, 4]));
assertEquals(5, maxFrequencyElements([1, 2, 3, 4, 5]));
