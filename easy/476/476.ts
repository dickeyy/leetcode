import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 46ms (beats 95.08% of TS solutions)
 *   Memory: 51.56mb (beats 16.39% of TS solutions)
 */

function findComplement(num: number): number {
    let binaryNum = num.toString(2).split("");
    let complement = "";
    binaryNum.forEach((dig) => {
        if (dig === "0") complement = complement + "1";
        else complement = complement + "0";
    });
    return parseInt(complement, 2);
}

// test cases
assertEquals(2, findComplement(5));
