import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 68ms (beats 56.44% of TS solutions)
 *   Memory: 54.16mb (beats 33.99% of TS solutions)
 */

function isPalindrome(s: string): boolean {
    const clean = s.replace(/[^a-zA-Z0-9]/g, "").toLowerCase();
    return clean === clean.split("").reverse().join("");
}

// test cases
assertEquals(true, isPalindrome("A man, a plan, a canal: Panama"));
assertEquals(false, isPalindrome("race a car"));
