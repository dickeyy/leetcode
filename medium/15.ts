/*
 *   Stats:
 *   Runtime: 152ms (beats 84.61% of TS solutions)
 *   Memory: 65.36mb (beats 69.73% of TS solutions)
 */

function threeSum(nums: number[]): number[][] {
    let result: number[][] = [];

    // return all triplets such that
    // a + b + c = 0
    // a, b, and c are unique
    // no duplicate triplets

    // first sort the nums array
    nums.sort((a, b) => a - b);

    // iterate through the nums array
    for (let i = 0; i < nums.length; i++) {
        // if the current number is greater than 0, then we can't have a sum of 0
        if (nums[i] > 0) {
            break;
        }

        // if the current number is the same as the previous number, then we've already checked it
        if (i > 0 && nums[i] === nums[i - 1]) {
            continue;
        }

        let left = i + 1;
        let right = nums.length - 1;

        while (left < right) {
            let sum = nums[i] + nums[left] + nums[right];

            if (sum === 0) {
                result.push([nums[i], nums[left], nums[right]]);

                while (nums[left] === nums[left + 1]) {
                    left++;
                }

                while (nums[right] === nums[right - 1]) {
                    right--;
                }

                left++;
                right--;
            } else if (sum < 0) {
                left++;
            } else {
                right--;
            }
        }
    }

    return result;
}

console.log(threeSum([-1, 0, 1, 2, -1, -4])); // [[-1,-1,2],[-1,0,1]]
console.log(threeSum([0, 1, 1])); // []
console.log(threeSum([0, 0, 0])); // [[0,0,0]]
