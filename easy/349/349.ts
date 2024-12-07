import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 52ms (beats 87.56% of TS solutions)
 *   Memory: 51.53mb (beats 61.96% of TS solutions)
 */

function intersection(nums1: number[], nums2: number[]): number[] {
    const arr: number[] = [];
    nums1.forEach((num) => {
        if (nums2.includes(num) && !arr.includes(num)) {
            arr.push(num);
        }
    });
    return arr;
}

// test cases
assertEquals([2], intersection([1, 2, 2, 1], [2, 2]));
assertEquals([4, 9], intersection([4, 9, 5], [9, 4, 9, 8, 4]));
