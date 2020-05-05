import typing
import itertools
from string import digits

class Luhn:
    def __init__(self, card_num: str) -> None:
        self.num = card_num.replace(' ','')


    '''for fun
    def valid(self) -> bool:
        """Confirm whether or not the provided card number is a valid number"""
        return (self.num.isdigit()) and len(self.num) != 1 and\
                (sum(itertools.chain([int(val) for val in self.num[::-2]],
                 [(dub := int(val) * 2) - (0 if dub <= 9 else 9) for val in self.num[-2::-2]])) % 10) == 0
    '''

    # for readability
    def valid(self) -> bool:
        """Confirm whether or not the provided card number is a valid number"""
        if not self.num.isdigit() or len(self.num) == 1:
            return False
        same_vals = [int(i) for i in self.num[::-2]]
        doubled = [int(val) * 2 for val in self.num[-2::-2]]
        combined = [num - (0 if num <= 9 else 9) for num in doubled]
        combined.extend(same_vals)
        return (sum(combined) % 10) == 0