import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 67ms (beats 78.33% of TS solutions)
 *   Memory: 59.30mb (beats 58.33% of TS solutions)
 */

function getCommon(nums1: number[], nums2: number[]): number {
    let i = 0;
    let j = 0;

    let common = -1;

    while (i < nums1.length && j < nums2.length) {
        if (nums1[i] === nums2[j]) {
            common = nums1[i];
            i++;
            j++;
        } else if (nums1[i] < nums2[j]) {
            i++;
        } else {
            j++;
        }

        if (common !== -1) {
            break;
        }
    }

    return common;
}

// test cases
assertEquals(2, getCommon([1, 2, 3], [2, 4]));
assertEquals(2, getCommon([1, 2, 3, 6], [2, 3, 4, 5]));
