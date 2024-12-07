/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
    const numsLength = nums.length;

    for (let i = 0; i < numsLength; i++) {
        for (let j = i + 1; j < numsLength; j++) {
            if (nums[i] + nums[j] === target) {
                return [i, j];
            }
        }
    }
};

// Works -- Submission accepted