/*
 *   Stats:
 *   Runtime: 2ms (beats 71.68% of TS solutions)
 *   Memory: 58.47mb (beats 96.68% of TS solutions)
 */

class TreeNode {
    val: number;
    left: TreeNode | null;
    right: TreeNode | null;
    constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
        this.val = val === undefined ? 0 : val;
        this.left = left === undefined ? null : left;
        this.right = right === undefined ? null : right;
    }
}

function sortedArrayToBST(nums: number[]): TreeNode | null {
    if (!nums) return null;

    function recurse(l: number, r: number): TreeNode | null {
        if (l > r) return null;
        const m = Math.floor((l + r) / 2);
        const node = new TreeNode(nums[m]);
        node.left = recurse(l, m - 1);
        node.right = recurse(m + 1, r);
        return node;
    }

    return recurse(0, nums.length - 1);
}
