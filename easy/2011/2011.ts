import assertEquals from "../../utils/assertEquals";

function finalValueAfterOperations(operations: string[]): number {
    let x = 0;
    for (let i = 0; i < operations.length; i++) {
        switch (operations[i]) {
            case "--X":
                x--;
                break;
            case "X--":
                x--;
                break;
            case "++X":
                x++;
                break;
            case "X++":
                x++;
                break;
        }
    }

    return x;
}

assertEquals(1, finalValueAfterOperations(["X++", "X++", "X--", "++X", "--X"]));
