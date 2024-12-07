// UNSOLVED UNSOLVED UNSOLVED UNSOLVED UNSOLVED UNSOLVED

/*
*   Stats:
*   Runtime: ms (beats % of TS solutions)
*   Memory: mb (beats % of TS solutions)
*/

function distance(nums: number[]): number[] {

    let arr: number[] = []

    for (let i = 0; i < nums.length; i++) {

        let num = nums[i]
        let eqiv: number[] = [] // array of indexes of nums whose values are equal to that of num, inclusive

        nums.forEach((value, index) => {
            if (value == num) {
                eqiv.push(index)
            }
        })

        let total = 0
        eqiv.forEach((value) => {
            total += Math.abs(i - value)
        })

        arr.push(total)

    }

    return arr

};

console.log(distance([1,3,1,1,2]))
// console.log(distance([0,5,3]))

// passes these test cases but runs into time limit exceed on large data sets... ill come back to this (probably)