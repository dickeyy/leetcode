/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function(nums) {
    
    var arr = []
    
    for (var i = 0; i < nums.length; i) {
        var num = nums[i]
        
        if (num == 0) {
            
            nums.splice(i, 1)
            arr.push(num)
            
        } else {
            i++
        }
        
    }

    for (var i = 0; i < arr.length; i++) {
        nums.push(arr[i])
    }
    
    
    return nums
};