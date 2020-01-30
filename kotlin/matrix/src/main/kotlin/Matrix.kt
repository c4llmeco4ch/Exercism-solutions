class Matrix(private val matrixAsString: String) {

    fun column(colNr: Int): List<Int> {
        var grid = matrixAsString.split("\n")
        var colInt: MutableList<Int> = mutableListOf<Int>()
        for (row in grid){ // See below note
            var nums = row.split(" ").filter{it -> it != ""} //Filter needed to deal with 
            if(nums.size < colNr){
                continue
            }
            colInt.add(nums[colNr - 1].toInt())
        }
        return colInt
    }

    fun row(rowNr: Int): List<Int> {
        var rowStr: List<String> = matrixAsString.split("\n")[rowNr - 1].split(" ")
        var rowInt: MutableList<Int> = mutableListOf<Int>()
        for(num in rowStr){ //Ideally I'd do this using functions without the loop
            if(num.length == 0){
                continue
            }
            rowInt.add(num.toInt())
        }
        return rowInt
        //Not sure how to fix this: return rowStr.asIntArray().toList()
    }
}
