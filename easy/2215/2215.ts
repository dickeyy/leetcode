import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 121ms (beats 18.05% of TS solutions) yikes...
 *   Memory: 56.76mb (beats 75.94% of TS solutions)
 */

// this is slow...
function findDifference(nums1: number[], nums2: number[]): number[][] {
    let answer: number[][] = [];

    answer.push(Array.from(new Set(nums1.filter((num) => !nums2.includes(num)))));
    answer.push(Array.from(new Set(nums2.filter((num) => !nums1.includes(num)))));

    return answer;
}

/*
 *   Stats:
 *   Runtime: 76ms (beats 88.13% of TS solutions) yikes...
 *   Memory: 57.32mb (beats 33.98% of TS solutions)
 */

// heres another version
function findDifference2(nums1: number[], nums2: number[]): number[][] {
    let set1 = new Set(nums1);
    let set2 = new Set(nums2);

    return [Array.from(set1).filter((n) => !set2.has(n)), Array.from(set2).filter((n) => !set1.has(n))];
}

// test cases
assertEquals(
    [
        [1, 3],
        [4, 6]
    ],
    findDifference([1, 2, 3], [2, 4, 6])
);
