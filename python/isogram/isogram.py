from string import ascii_letters
def is_isogram(string):
    return (s:=string.lower()) == '' or len([_ for _, val in enumerate(s) if val in ascii_letters and s.count(val) != 1]) == 0
