package leedcodeproblm

type Solution struct {
	memo  map[int]int
	memo1 map[rune]int
}

func (s *Solution) NumDecodings(str string) int {
	s.memo = make(map[int]int)
	sum := s.numberOfDecode(str, 0)
	return sum
}

func (s *Solution) numberOfDecode(str string, index int) int {
	if val, ok := s.memo[index]; ok {
		return val
	}
	if len(str) > index && str[index] == '0' {
		return 0
	}
	if index == len(str) {
		return 1
	}
	if len(str)-1 == index {
		return 1
	}
	left := s.numberOfDecode(str, index+1)
	right := 0
	if index < len(str)-1 && (str[index] == '1' || (str[index] == '2' && str[index+1] <= '6')) {
		right = s.numberOfDecode(str, index+2)
	}
	s.memo[index] = right + left
	return right + left
}
