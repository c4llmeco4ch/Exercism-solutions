class Matrix:
    def __init__(self, matrix_string):
        self.matrix = []
        if matrix_string == '':
            return
        rows = matrix_string.split('\n')
        self.matrix = [[int(col) for col in row.split()] for row in rows]

    def row(self, index):
        return self.matrix[index - 1]

    def column(self, index):
        return [self.matrix[i][index - 1] for i,_ in enumerate(self.matrix)]
