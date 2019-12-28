def convert(number):
    strings = ['Pling' if not number % 3 else '',
               'Plang' if not number % 5 else '',
               'Plong' if not number % 7 else '']
    answer = ''.join(strings)
    return answer if len(answer) > 0 else str(number)
