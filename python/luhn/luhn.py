import typing
import itertools
from string import digits

class Luhn:
    def __init__(self, card_num: str) -> None:
        self.num = [i for i in card_num if i != " "]


    '''for fun
    def valid(self) -> bool:
        """Confirm whether or not the provided card number is a valid number"""
        return (not any(i for i in self.num if i not in digits)) and\
                len(self.num) != 1 and\
                (sum(itertools.chain([int(val) for val in self.num[::-2]],
                 [(dub := int(val) * 2) - (0 if dub <= 9 else 9) for val in self.num[-2::-2]])) % 10) == 0
    '''

    # for readability
    def valid(self) -> bool:
        """Confirm whether or not the provided card number is a valid number"""
        if any(i for i in self.num if i not in digits):
            return False
        same_vals = [int(i) for i in self.num[::-2]]
        doubled = [int(val) * 2 for val in self.num[-2::-2]]
        for pos, num in enumerate(doubled):
            if num > 9:
                doubled[pos] = num - 9
        same_vals.extend(doubled)
        return len(same_vals) != 1 and (sum(same_vals) % 10) == 0