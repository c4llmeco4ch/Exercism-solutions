PLANT_ALTS = {'R': 'Radishes', 'G': 'Grass', 'C': 'Clover', 'V': 'Violets'}


class Garden:
    """
    Given a selection of student names and a complementary diagram,
    create a classroom garden
    """
    def __init__(self, diagram, students=["Alice", "Bob", "Charlie",
                                          "David", "Eve", "Fred",
                                          "Ginny", "Harriet", "Ileana",
                                          "Joseph", "Kincaid", "Larry"]):
        self.diagram = diagram.splitlines()
        self.students = sorted(students)

    def plants(self, student):
        """Given a student, return a list of plants they are responsible for"""
        group = self.students.index(student) * 2
        return [PLANT_ALTS[plant] for plant in
                [plant for row in self.diagram for plant in row[group:group+2]]]
