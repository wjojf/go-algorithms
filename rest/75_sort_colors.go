package main

func sortColors(n []int)  {
    var i = 1
    for i < len(n) {
        var j = i
        for j >= 1 && n[j] < n[j - 1] {
            n[j], n[j - 1] = n[j - 1], n[j]

            j--
        }

        i++
    }
}
