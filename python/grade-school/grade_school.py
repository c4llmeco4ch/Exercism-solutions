class School:
    def __init__(self):
        self.all_grades = {}

    def add_student(self, name, grade):
        """
        given a student and grade,
        add that student to the grade's roster
        """
        if grade not in self.all_grades:
            self.all_grades[grade] = []
        self.all_grades[grade].append(name)

    def roster(self):
        """return the roster sorted by grade then by name"""
        return [student for g in sorted(self.all_grades.keys())
                for student in self.grade(g)]

    def grade(self, grade_number):
        """return a sorted list of the students in a grade"""
        return sorted(self.all_grades[grade_number])\
            if grade_number in self.all_grades else []
