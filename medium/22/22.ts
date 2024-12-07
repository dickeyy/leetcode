import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 59ms (beats 58.35% of TS solutions)
 *   Memory: 51.74mb (beats 46.17% of TS solutions)
 */

function generateParenthesis(n: number): string[] {
    const result: string[] = [];
    function backtrack(s: string, open: number, close: number) {
        if (s.length === n * 2) {
            result.push(s);
            return;
        }
        if (open < n) backtrack(s + "(", open + 1, close);
        if (close < open) backtrack(s + ")", open, close + 1);
    }
    backtrack("", 0, 0);
    return result;
}

// Test cases
assertEquals(["((()))", "(()())", "(())()", "()(())", "()()()"], generateParenthesis(3));
assertEquals(["()"], generateParenthesis(1));
