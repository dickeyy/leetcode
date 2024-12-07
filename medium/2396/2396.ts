import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 59ms (beats 41.56% of TS solutions)
 *   Memory: 49.62mb (beats 66.23% of TS solutions)
 */

function isStrictlyPalindromic(n: number): boolean {
    // this function will ALWAYS return false based of the nature of the question.
    // The number n in base (n - 2) is always 12, which is not palindromic, therefore, it will always return false
    return false;
}

// test cases
assertEquals(false, isStrictlyPalindromic(9));
assertEquals(false, isStrictlyPalindromic(4));
