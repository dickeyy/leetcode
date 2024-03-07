import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 55ms (beats 70.97% of TS solutions)
 *   Memory: 52.60mb (beats 16.13% of TS solutions)
 */

function prefixCount(words: string[], pref: string): number {
    let count = 0;
    words.forEach((word) => {
        if (word.startsWith(pref)) count++;
    });
    return count;
}

// Test cases
assertEquals(2, prefixCount(["pay", "attention", "practice", "attend"], "at"));
assertEquals(0, prefixCount(["leetcode", "win", "loops", "success"], "code"));
