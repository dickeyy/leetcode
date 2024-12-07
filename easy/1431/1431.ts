import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 50ms (beats 96.12% of TS solutions)
 *   Memory: 52.92mb (beats 19.78% of TS solutions)
 */

function kidsWithCandies(candies: number[], extraCandies: number): boolean[] {
    let arr = Array(candies.length).fill(false);
    let max = [...candies].sort((a, b) => b - a)[0];

    for (let i = 0; i < candies.length; i++) {
        if (candies[i] + extraCandies >= max) arr[i] = true;
    }

    return arr;
}

// test cases
assertEquals([true, true, true, false, true], kidsWithCandies([2, 3, 5, 1, 3], 3));
