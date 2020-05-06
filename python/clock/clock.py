class Clock:

    HOURS_IN_DAY = 24

    def __init__(self, hour, minute):
        self.hour = hour
        self.minute = minute
        self.reduce_minutes()


    def __repr__(self):
        ''' Is there a way to enforce 0s with an f-string?
        return f'{self.hour}:{self.minute}'
        '''
        return '%02d:%02d' % (self.hour, self.minute)

    def __eq__(self, other):
        return self.minute == other.minute and self.hour == other.hour

    def __add__(self, minutes):
        self.minute += minutes
        self.reduce_minutes()
        return self

    def __sub__(self, minutes):
        self.minute -= minutes
        self.reduce_minutes()
        return self

    def reduce_minutes(self):
        while self.minute >= 60:
            self.minute -= 60
            self.hour += 1
        while self.minute < 0:
            self.minute += 60
            self.hour -= 1
        self.hour %= self.HOURS_IN_DAY