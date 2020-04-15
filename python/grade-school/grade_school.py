from collections import defaultdict
import typing
from typing import List
class School:
    def __init__(self) -> None: 
        self.all_grades = defaultdict(list)

    def add_student(self, name: str, grade: int) -> None:
        """
        given a student and grade,
        add that student to the grade's roster
        """
        self.all_grades[grade].append(name)

    def roster(self) -> List[List[str]]:
        """return the roster sorted by grade then by name"""
        return [student for g in sorted(self.all_grades.keys())
                for student in self.grade(g)]

    def grade(self, grade_number: int) -> List[str]:
        """return a sorted list of the students in a grade"""
        return sorted(self.all_grades[grade_number])\
            if grade_number in self.all_grades else []
