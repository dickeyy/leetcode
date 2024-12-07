import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 47ms (beats 95.88% of TS solutions)
 *   Memory: 51.80mb (beats 18.04% of TS solutions)
 */

function maximumWealth(accounts: number[][]): number {
    let max = 0;
    for (let i = 0; i < accounts.length; i++) {
        let total = 0;
        for (let j = 0; j < accounts[i].length; j++) {
            total += accounts[i][j];
        }
        if (total > max) max = total;
    }
    return max;
}

// Test cases
assertEquals(
    6,
    maximumWealth([
        [1, 2, 3],
        [3, 2, 1]
    ])
);
assertEquals(
    10,
    maximumWealth([
        [1, 5],
        [7, 3],
        [3, 5]
    ])
);
