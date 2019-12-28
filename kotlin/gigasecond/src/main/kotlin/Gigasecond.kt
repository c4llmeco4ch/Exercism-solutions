import java.time.LocalDateTime
import java.time.LocalDate

class Gigasecond(_date: LocalDateTime){
    val date: LocalDateTime

    init{
        date = _date.bumpForward()
    }

    constructor(_date: LocalDate): this(_date.atStartOfDay()){
    }

    fun LocalDateTime.bumpForward(): LocalDateTime {
        return this.plusSeconds(1_000_000_000)
    }
}
