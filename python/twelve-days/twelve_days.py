days = [['first', 'a Partridge in a Pear Tree'],
        ['second', 'two Turtle Doves'],
        ['third', 'three French Hens'],
        ['fourth', 'four Calling Birds'],
        ['fifth', 'five Gold Rings'],
        ['sixth', 'six Geese-a-Laying'],
        ['seventh', 'seven Swans-a-Swimming'],
        ['eighth', 'eight Maids-a-Milking'],
        ['ninth', 'nine Ladies Dancing'],
        ['tenth', 'ten Lords-a-Leaping'],
        ['eleventh', 'eleven Pipers Piping'],
        ['twelfth', 'twelve Drummers Drumming'],
        ]


def recite(start_verse, end_verse):
    return [calc_day(day) for day in range(start_verse - 1, end_verse)]


def calc_day(day):
    text = ('On the %s day of Christmas ' +
            'my true love gave to me: ') % days[day][0]
    for i in range(day, -1, -1):
        text = ''.join([text,
                        (', ' if i != day else ''),
                        ('and ' if day != 0 and i == 0 else ''),
                        days[i][1]])
    return ''.join([text, '.'])
