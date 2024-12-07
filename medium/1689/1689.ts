import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 64ms (beats 90.00% of TS solutions)
 *   Memory: 54.78mb (beats 64.00% of TS solutions)
 */

// First attempt (works but is super inefficient)
// function minPartitions(n: string): number {
//     let nums: number[] = [];
//     n.split("").forEach((num) => {
//         nums.push(Number(num));
//     });
//     return nums.sort((a, b) => b - a)[0];
// }

// Second attempt (this one is much better)
function minPartitions(n: string): number {
    let max = 0;
    for (let i = 0; i < n.length; i++) {
        if (+n[i] > max) max = +n[i];
    }
    return max;
}

// test cases
assertEquals(3, minPartitions("32"));
assertEquals(8, minPartitions("82734"));
assertEquals(9, minPartitions("27346209830709182346"));
