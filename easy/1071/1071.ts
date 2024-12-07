import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 50ms (beats 92.57% of TS solutions)
 *   Memory: 51.50mb (beats 69.11% of TS solutions)
 */

function gcdOfStrings(str1: string, str2: string): string {
    if (str1 + str2 !== str2 + str1) return "";
    return str1.substring(0, gcd(str1.length, str2.length));
}

function gcd(a: number, b: number): number {
    if (b === 0) return a;
    return gcd(b, a % b);
}

// test cases
assertEquals("ABC", gcdOfStrings("ABCABC", "ABC"));
assertEquals("AB", gcdOfStrings("ABABAB", "ABAB"));
