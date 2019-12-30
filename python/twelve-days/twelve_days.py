gifts = ['a Partridge in a Pear Tree',
         'two Turtle Doves',
         'three French Hens',
         'four Calling Birds',
         'five Gold Rings',
         'six Geese-a-Laying' ,
         'seven Swans-a-Swimming',
         'eight Maids-a-Milking',
         'nine Ladies Dancing',
         'ten Lords-a-Leaping',
         'eleven Pipers Piping',
         'twelve Drummers Drumming',
        ]

days = ['first', 'second', 'third', 'fourth', 
        'fifth', 'sixth', 'seventh', 'eighth',
        'ninth', 'tenth', 'eleventh', 'twelfth']


def recite(start_verse, end_verse):
    return [calc_day(day) for day in range(start_verse - 1, end_verse)]


def calc_day(day):
    text = ('On the %s day of Christmas ' +
            'my true love gave to me: ') % days[day]
    all_gifts = ', '.join([gifts[i] for i in range(day, 0, -1)])
    return ''.join([text, all_gifts, 
                    ', and ' if day > 0 else '', 
                    gifts[0], '.'])
