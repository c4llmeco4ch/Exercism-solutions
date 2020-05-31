class Clock:

    HOURS_IN_DAY = 24
    MINUTES_IN_HOUR = 60

    def __init__(self, hour, minute):
        self.hour = hour
        self.minute = minute
        self.reduce_minutes()

    def __repr__(self):
        return f'{self.hour:02}:{self.minute:02}'

    def __eq__(self, other):
        return self.minute == other.minute and self.hour == other.hour

    def __add__(self, minutes):
        return self.__class__(self.hour, self.minute + minutes)

    def __sub__(self, minutes):
        return self.__class__(self.hour, self.minute - minutes)

    def reduce_minutes(self):
        self.hour = (self.hour + (self.minute // 60)) % self.HOURS_IN_DAY
        self.minute %= self.MINUTES_IN_HOUR
