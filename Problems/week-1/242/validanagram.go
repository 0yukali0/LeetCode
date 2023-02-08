func isAnagram(s string, t string) bool {
    a := make(map[rune]uint64, 0)
    b := make(map[rune]uint64, 0)
    for _, c := range s {
        if value, ok := a[c]; ok {
            a[c] = value + 1
        } else {
            a[c] = 1
        }
    }
    for _, c := range t {
        if value, ok := b[c]; ok {
            b[c] = value + 1
        } else {
            b[c] = 1
        }
    }

    if len(a) != len(b) {
        return false
    } else {
        for key, expect := range a {
            if got, ok := b[key]; !ok || got != expect {
                return false
            }
        }
    }
    return true
}
