import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 63ms (beats 32.16% of TS solutions)
 *   Memory: 51.85mb (beats 50.16% of TS solutions)
 */

function mergeAlternately(word1: string, word2: string): string {
    let merge = "";
    const minL = Math.min(word1.length, word2.length);

    for (let i = 0; i < minL; i++) {
        merge += word1[i] + word2[i];
    }

    merge += word1.slice(minL) + word2.slice(minL);

    return merge;
}

// test cases
assertEquals("apbqcr", mergeAlternately("abc", "pqr"));
assertEquals("apbqrs", mergeAlternately("ab", "pqrs"));
assertEquals("apbqcd", mergeAlternately("abcd", "pq"));
