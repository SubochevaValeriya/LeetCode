class Solution:
    def dailyTemperatures(self, temperatures):
        answer = [0] * len(temperatures)
        stack = []
        for day, temp in enumerate(temperatures):
            print(day, temp)
            while stack and temperatures[stack[-1]] < temp:
                index = stack.pop()
                answer[index] = day - index
            stack.append(day)
        return answer

print(Solution.dailyTemperatures(1, [73, 74, 75, 71, 69, 72, 76, 73]))
