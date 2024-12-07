import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 58ms (beats 59.01% of TS solutions)
 *   Memory: 51.11mb (beats 67.93% of TS solutions)
 */

// didnt read the whole problem first and originally just did return Math.pow(x, n); lol
// heres the actual solution implementing my own pow function (recursion (yuck))

function myPow(x: number, n: number): number {
    if (n === 0) return 1;
    if (n < 0) {
        x = 1 / x;
        n = -n;
    }
    return n % 2 === 0 ? myPow(x * x, n / 2) : x * myPow(x * x, Math.floor(n / 2));
}

// test cases
assertEquals(1024, myPow(2, 10));
assertEquals(9.261, myPow(2.1, 3));
assertEquals(0.25, myPow(2, -2));
