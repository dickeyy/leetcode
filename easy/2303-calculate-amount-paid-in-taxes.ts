function calculateTax(brackets: number[][], income: number) {
    if (income <= 0) return 0; // make no money, dont have to pay taxes :)

    let total: number = 0;
    let prev: number = 0; // the previous tax brackets upper bound
    let percentages: number[] = [];
    let dollars: number[] = [];

    // for each tax bracket (bracket[i])
    for (let i = 0; i < brackets.length; i++) {
        if (income > brackets[i][0]) continue;

        percentages.push(brackets[i][1] / 100);

        if (prev <= 0) dollars.push(brackets[i][0]);
        else {
            if (brackets[i][0] < income) dollars.push(brackets[i][0] - prev);
            else dollars.push(income - prev);
        }

        prev = brackets[i][0];
    }

    for (let i = 0; i < dollars.length; i++) {
        total += dollars[i] * percentages[i];
    }

    return total;
}

console.log(
    calculateTax(
        [
            [1, 0],
            [4, 25],
            [5, 50]
        ],
        2
    )
);
