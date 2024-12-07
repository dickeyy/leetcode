import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 76ms (beats 70.94% of TS solutions)
 *   Memory: 54.06mb (beats 27.36% of TS solutions)
 */

function findWordsContaining(words: string[], x: string): number[] {
    let arr: number[] = [];
    words.forEach((word, i) => {
        if (word.includes(x)) {
            arr.push(i);
        }
    });

    return arr;
}

// Test cases
assertEquals([0, 1], findWordsContaining(["leet", "code"], "e"));
