/*
*   Stats:
*   Runtime: 57ms (beats 52.40% of TS solutions)
*   Memory: 44.18mb (beats 42.78% of TS solutions)
*/

function removeElement(nums: number[], val: number): number {

    let start = 0;
    let end = nums.length - 1;

    let k = nums.length;

    while(start <= end) {
        if (nums[start] === val) {
            if (nums[end] === val) {
                end --;
                k --;
                continue
            } else {
                nums[start] = nums[end]
                nums[end] = -1
                end --
                k --
            }
        }
        start ++
    }

    return k;

};

console.log(removeElement([0,1,2,2,3,0,4,2],2))

// this is a horribly written question.