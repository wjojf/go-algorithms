package main

func romanToInt(s string) int {
    var RomanNumerals = map[rune]int{
    	'I': 1,
    	'V': 5,
    	'X': 10,
    	'L': 50,
	    'C': 100,
	    'D': 500,
	    'M': 1000,
    }

    var sum int 
    var greatest int 

    for i := len(s) - 1; i >= 0; i--{
        letter := s[i]
        num := RomanNumerals[rune(letter)]
        if num >= greatest {
            greatest = num
            sum += num 
            continue 
        }

        sum -= num

    }

    return sum
}
