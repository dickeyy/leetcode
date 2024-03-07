import assertEquals from "../utils/assertEquals";

/*
 *   Stats:
 *   Runtime: 46ms (beats 95.25% of TS solutions)
 *   Memory: 51.34mb (beats 37.60% of TS solutions)
 */

function middleNode(head: any | null): any | null {
    // given head of a singly linekd list, return the middle node
    // if there are two middle nodes, return the second one
    let slow = head;
    let fast = head;
    while (fast && fast.next) {
        slow = slow.next;
        fast = fast.next.next;
    }
    return slow;
}

// Test cases
assertEquals([3, 4, 5], middleNode([1, 2, 3, 4, 5]));
assertEquals([4, 5, 6], middleNode([1, 2, 3, 4, 5, 6]));
