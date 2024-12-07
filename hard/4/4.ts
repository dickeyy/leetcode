import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 102ms (beats 54.63% of TS solutions) :)
 *   Memory: 56.50mb (beats 51.64% of TS solutions) :(
 */

// this works, but it doesnt meet the time complexity requirement of O(log(m+n)) (i don't think)
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
    let merged = [...nums1, ...nums2].sort((a, b) => a - b);
    let middleIndex = Math.floor(merged.length / 2);

    if (merged.length % 2 === 1) return merged[middleIndex];

    return (merged[middleIndex - 1] + merged[middleIndex]) / 2;
}

// test cases
assertEquals(2, findMedianSortedArrays([1, 3], [2]));
assertEquals(2.5, findMedianSortedArrays([1, 2], [3, 4]));
