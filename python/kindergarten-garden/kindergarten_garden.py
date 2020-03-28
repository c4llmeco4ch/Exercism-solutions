import itertools

PLANT_ALTS = {'R' : 'Radishes', 'G' : 'Grass', 'C' : 'Clover', 'V' : 'Violets'}
class Garden:
    """Given a selection of student names and a complementary diagram, create a classroom garden"""
    def __init__(self, diagram, students = ["Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"]):
        self.diagram = diagram.split('\n')
        self.students = sorted(students)
    
    def plants(self, student):
        """Given a student, return a list of plants they are responsible for"""
        group = self.students.index(student) * 2
        vals = list(itertools.chain(self.diagram[0][group:group + 2], self.diagram[1][group:group + 2]))
        for pos, plant in enumerate(vals):
            vals[pos] = PLANT_ALTS[plant]
        return vals

