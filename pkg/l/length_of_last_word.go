package l

func lengthOfLastWord(s string) int {
    res := 0
    trim := true
    for i := len(s) - 1; i >= 0; i-- {
        if s[i:i+1] == " " && trim {
            continue
        } else if s[i : i + 1] != " " {
            trim = false
            res ++
        } else {
            break
        }
    }

    return res

}
