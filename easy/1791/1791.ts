import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 92ms (beats 74.47% of TS solutions)
 *   Memory: 66.38mb (beats 61.70% of TS solutions)
 */

function findCenter(edges: number[][]): number {
    return edges[1].includes(edges[0][0]) ? edges[0][0] : edges[0][1];
}

// test cases
assertEquals(
    2,
    findCenter([
        [1, 2],
        [2, 3],
        [4, 2]
    ])
);
assertEquals(
    1,
    findCenter([
        [1, 2],
        [5, 1],
        [1, 3],
        [1, 4]
    ])
);
