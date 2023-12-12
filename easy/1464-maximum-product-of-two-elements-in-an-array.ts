/*
 *   Stats:
 *   Runtime: 67ms (beats 23.23% of TS solutions)
 *   Memory: 44.61mb (beats 67.68% of TS solutions)
 */

function maxProduct(nums: number[]): number {
    const sorted = nums.sort((a, b) => b - a);
    return (sorted[0] - 1) * (sorted[1] - 1);
}

console.log(maxProduct([3, 4, 5, 2]));
console.log(maxProduct([1, 5, 4, 5]));
console.log(maxProduct([3, 7]));
