/*
 *   Stats:
 *   Runtime: 56ms (beats 90.70% of TS solutions)
 *   Memory: 46.44mb (beats 72.09% of TS solutions)
 */

function largestOddNumber(num: string): string {
    for (let i = num.length - 1; i >= 0; i--) {
        if (parseInt(num[i]) % 2 !== 0) {
            return num.substring(0, i + 1);
        }
    }

    return "";
}

console.log(largestOddNumber("4206"));
