class Solution(object):
    def finalValueAfterOperations(self, operations):
        """
        :type operations: List[str]
        :rtype: int
        """
        
        x = 0
        for operation in operations:
            if operation == "--X":
                x -= 1
            elif operation == "X--":
                x -= 1
            elif operation == "++X":
                x += 1
            elif operation == "X++":
                x += 1
        return x    

print(Solution().finalValueAfterOperations(["X++", "X++", "X--", "++X", "--X"]) == 1)