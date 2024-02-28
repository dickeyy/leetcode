/*
 *   Stats:
 *   Runtime: 112ms (beats 73.94% of TS solutions)
 *   Memory: 57.48mb (beats 26.67% of TS solutions)
 */

function intToRoman(num: number): string {
    let str = "";

    const roman = [["I", "V"], ["X", "L"], ["C", "D"], ["M"]];

    let i = 0;
    while (num > 0) {
        let digit = num % 10;
        num = Math.floor(num / 10);

        if (digit === 9) {
            str = roman[i][0] + roman[i + 1][0] + str;
        } else if (digit >= 5) {
            str = roman[i][1] + roman[i][0].repeat(digit - 5) + str;
        } else if (digit === 4) {
            str = roman[i][0] + roman[i][1] + str;
        } else {
            str = roman[i][0].repeat(digit) + str;
        }

        i++;
    }

    return str;
}

console.log(intToRoman(3)); // "III"
console.log(intToRoman(58)); // "LVIII"
console.log(intToRoman(1994)); // "MCMXCIV"
