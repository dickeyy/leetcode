/*
 *   Stats:
 *   Runtime: 72ms (beats 5.00% of TS solutions)
 *   Memory: 43.22mb (beats 30.00% of TS solutions)
 */

function totalMoney(n: number): number {
    const weeks = Math.floor(n / 7);
    let total = 0;

    // add the full weeks
    for (let i = 0; i < weeks; i++) {
        total += 28 + i * 7;
    }

    // add the remaining days
    for (let i = 0; i < n % 7; i++) {
        total += weeks + i + 1;
    }

    return total;
}

console.log(totalMoney(4));
