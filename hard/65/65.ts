import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 85ms (beats 54.27% of TS solutions)
 *   Memory: 53.98mb (beats 70.12% of TS solutions)
 */

// i have a feeling, given that this problem has a solution rate of 19%, that this is not the intended solution
// that being said, it passes all the cases sooooo
function isNumber(s: string): boolean {
    if (s === "Infinity" || s === "-Infinity" || s === "+Infinity") return false;
    return !isNaN(Number(s));
}

// Test cases
assertEquals(true, isNumber("0"));
