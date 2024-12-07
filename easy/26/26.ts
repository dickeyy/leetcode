/*
 *   Stats:
 *   Runtime: 67ms (beats 65.26% of TS solutions)
 *   Memory: 52.98mb (beats 44.58% of TS solutions)
 */

function removeDuplicates(nums: number[]): number {
    let temp = [...new Set(nums)];
    nums.length = 0;
    nums.push(...temp);
    return nums.length;
}

// test cases
console.log(removeDuplicates([1, 1, 2])); // 2
console.log(removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4])); // 5
console.log(removeDuplicates([1, 1, 2])); // 2
