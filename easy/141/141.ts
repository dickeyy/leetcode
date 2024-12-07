import assertEquals from "../../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 70ms (beats 55.54% of TS solutions)
 *   Memory: 54.22mb (beats 55.54% of TS solutions)
 */

function hasCycle(head: any | null): boolean {
    if (!head) return false;
    let slow = head;
    let fast = head.next;
    while (slow !== fast) {
        if (!fast || !fast.next) return false;
        slow = slow.next;
        fast = fast.next.next;
    }
    return true;
}

// test cases
assertEquals(true, hasCycle([3, 2, 0, -4]));
