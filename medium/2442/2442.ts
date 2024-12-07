import assertEquals from "../utils/assertEquals";

/*
 *   Stats (Attempt 1):
 *   Runtime: 360ms (beats 14.29% of TS solutions)
 *   Memory: 77.54mb (beats 85.71% of TS solutions)
 */

// this is pretty bad lol. i have a second attempt below that is much better.
function countDistinctIntegersBad(nums: number[]): number {
    let rev: number[] = [];

    for (let i = 0; i < nums.length; i++) {
        if (nums[i] >= 10) {
            let revDigits: string[] = [];
            let split = nums[i].toString().split("");
            split.forEach((dig) => {
                revDigits.unshift(dig);
            });
            rev.push(Number(revDigits.join("")));
        } else {
            rev.push(nums[i]);
        }
    }

    return new Set(nums.concat(rev)).size;
}

// test cases
assertEquals(6, countDistinctIntegersBad([1, 13, 10, 12, 31]));
assertEquals(1, countDistinctIntegersBad([2, 2, 2]));

/*
 *   Stats (Attempt 2):
 *   Runtime: 255ms (beats 71.43% of TS solutions)
 *   Memory: 77.18mb (beats 92.86% of TS solutions)
 */

// okay attempt #2
function countDistinctIntegers(nums: number[]): number {
    const set = new Set(nums);
    for (let num of nums) {
        const digit = +String(num).split("").reverse().join("");
        set.add(digit);
    }
    return set.size;
}

// test cases
assertEquals(6, countDistinctIntegers([1, 13, 10, 12, 31]));
assertEquals(1, countDistinctIntegers([2, 2, 2]));
