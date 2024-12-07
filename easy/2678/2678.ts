import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 67ms (beats 25.96% of TS solutions)
 *   Memory: 52.21mb (beats 29.88% of TS solutions)
 */

// this solution sucks and is slow but it does work i guess
function countSeniors(details: string[]): number {
    let old = 0;

    for (let i in details) {
        const age = Number(
            details[i].slice(10, details[i].length).includes("M")
                ? details[i].slice(10, details[i].length).split("M")[1].slice(0, 2)
                : details[i].slice(10, details[i].length).includes("F")
                ? details[i].slice(10, details[i].length).split("F")[1].slice(0, 2)
                : details[i].slice(10, details[i].length).split("O")[1].slice(0, 2)
        );

        if (age > 60) {
            old++;
        }
    }

    return old;
}

// test cases
assertEquals(2, countSeniors(["7868190130M7522", "5303914400F9211", "9273338290F4010"]));
