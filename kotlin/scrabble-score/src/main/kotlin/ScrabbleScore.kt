object ScrabbleScore {

    fun scoreLetter(c: Char): Int {
        return when (c) {
            'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't' -> 1
            'd', 'g' -> 2
            'b', 'c', 'm', 'p' -> 3
            'f', 'h', 'v', 'w', 'y' -> 4
            'k' -> 5
            'j', 'x' -> 8
            'q', 'z' -> 10
            else -> -1
        }
    }

    fun scoreWord(word: String): Int {
        var sum: Int = 0
        for (c in word.toLowerCase()) {
            sum += scoreLetter(c)
        }
        return sum
    }
}
