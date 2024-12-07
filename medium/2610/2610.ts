import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 101ms (beats 78.48% of TS solutions)
 *   Memory: 58.00mb (beats 37.97% of TS solutions)
 */

function findMatrix(nums: number[]): number[][] {
    let r: number[][] = [];
    let m: Map<number, number> = new Map();

    for (let num of nums) {
        let i = m.get(num) || 0;
        if (r.length <= i) {
            r.push([]);
        }
        r[i].push(num);
        m.set(num, i + 1);
    }

    return r;
}

// test cases
assertEquals([[1, 3, 4, 2], [1, 3], [1]], findMatrix([1, 3, 4, 1, 2, 3, 1]));
