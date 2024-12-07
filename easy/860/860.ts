import assertEquals from "../../utils/assertEquals";

function lemonadeChange(bills: number[]): boolean {
    let change: number = 0;
    let canServe = true;

    if (bills[0] > 5) return false;

    for (let i = 0; i < bills.length; i++) {
        let denom = bills[i];

        console.log(`Customer: ${i}, Bill: ${denom}, Change: ${change}`);

        if (denom > 5) {
            const changeDue = denom - 5;
            if (changeDue > change) {
                // no change to give back to customer
                canServe = false;
                break;
            } else {
                // we have change for their bill
                change -= denom - 5; // give the customer back whatever the dif. is from change-bin
            }
        }

        change += denom;
    }

    return canServe;
}

// test cases
assertEquals(true, lemonadeChange([5, 5, 5, 10, 20]));
assertEquals(false, lemonadeChange([5, 5, 10, 10, 20]));
// 5 -> change=5
// 5 -> change=10
// 10 -> change=15
// 10 -> change=20
// 20 -> change=

// assertEquals(
//     true,
//     lemonadeChange([
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20,
//         5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 20
//     ])
// );
