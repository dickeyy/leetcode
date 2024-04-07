import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 63ms (beats 39.68% of TS solutions)
 *   Memory: 50.28mb (beats 87.30% of TS solutions)
 */

function checkValidString(s: string): boolean {
    let low = 0;
    let high = 0;

    for (let i = 0; i < s.length; i++) {
        if (s[i] === "(") {
            low++;
            high++;
        } else if (s[i] === ")") {
            low--;
            high--;
        } else {
            low--;
            high++;
        }

        if (high < 0) {
            return false;
        }

        if (low < 0) {
            low = 0;
        }
    }

    return low === 0;
}

// test cases
assertEquals(true, checkValidString("()"));
assertEquals(true, checkValidString("(*)"));
assertEquals(true, checkValidString("(*))"));
assertEquals(
    false,
    checkValidString(
        "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"
    )
);
