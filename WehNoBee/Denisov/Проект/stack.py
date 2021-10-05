class Stack:
    _stack = list()

    def push(self, elem):
        self._stack.append(elem)

    def pop(self):
        return self._stack.pop()

    def dump(self):
        return tuple(self._stack)