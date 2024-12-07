import assertEquals from "../../utils/assertEquals";

function twoSum(nums: number[], target: number): number[] {
    const numsLength = nums.length;

    for (let i = 0; i < numsLength; i++) {
        for (let j = i + 1; j < numsLength; j++) {
            if (nums[i] + nums[j] === target) {
                return [i, j];
            }
        }
    }

    return [];
}

assertEquals([0, 1], twoSum([2, 7, 11, 15], 9));
