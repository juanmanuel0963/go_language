package main

import "testing"

func Test_MaxStack_Max(t *testing.T) {
	s := MaxStack{}
	_, err := s.Max()
	if err == nil {
		t.Error("expected error")
	}
	s.Push(1)
	max, err := s.Max()
	if err != nil {
		t.Error("unexpected error")
	}
	if max != 1 {
		t.Error("expected 1")
	}
	s.Push(2)
	max, err = s.Max()
	if err != nil {
		t.Error("unexpected error")
	}
	if max != 2 {
		t.Error("expected 2")
	}
	s.Push(1)
	max, err = s.Max()
	if err != nil {
		t.Error("unexpected error")
	}
	if max != 2 {
		t.Error("expected 2")
	}
	s.Pop()
	max, err = s.Max()
	if err != nil {
		t.Error("unexpected error")
	}
	if max != 2 {
		t.Error("expected 2")
	}
	s.Pop()
	max, err = s.Max()
	if err != nil {
		t.Error("unexpected error")
	}
	if max != 1 {
		t.Error("expected 1")
	}
	s.Pop()
	_, err = s.Max()
	if err == nil {
		t.Error("expected error")
	}
}

func BenchmarkMaxStack_Max(b *testing.B) {
	s := MaxStack{}
	for i := 0; i < b.N; i++ {
		s.Max()
	}
}

func Test_MaxStack_Pop(t *testing.T) {
	s := MaxStack{}
	_, err := s.Pop()
	if err == nil {
		t.Error("expected error")
	}
	s.Push(1)
	x, err := s.Pop()
	if err != nil {
		t.Error("unexpected error")
	}
	if x != 1 {
		t.Error("expected 1")
	}
	_, err = s.Pop()
	if err == nil {
		t.Error("expected error")
	}
}

func BenchmarkMaxStack_Pop(b *testing.B) {
	s := MaxStack{}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func Test_MaxStack_Push(t *testing.T) {
	s := MaxStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	s.Push(9)
	s.Push(10)
	s.Push(11)
	s.Push(12)
	s.Push(13)
	s.Push(14)
	s.Push(15)
	s.Push(16)
	s.Push(17)
	s.Push(18)
	s.Push(19)
	s.Push(20)
	s.Push(21)
	s.Push(22)
	s.Push(23)
	s.Push(24)
	s.Push(25)
	s.Push(26)
	s.Push(27)
	s.Push(28)
	s.Push(29)
	s.Push(30)
	s.Push(31)
	s.Push(32)
	s.Push(33)
	s.Push(34)
	s.Push(35)
	s.Push(36)
	s.Push(37)
	s.Push(38)
	s.Push(39)
	s.Push(40)
	s.Push(41)
	s.Push(42)
	s.Push(43)
	s.Push(44)
	s.Push(45)
	s.Push(46)
	s.Push(47)
	s.Push(48)
	s.Push(49)
	s.Push(50)
	s.Push(51)
	s.Push(52)
	s.Push(53)
	s.Push(54)
	s.Push(55)
	s.Push(56)
	s.Push(57)
	s.Push(58)
	s.Push(59)
	s.Push(60)
	s.Push(61)
	s.Push(62)
	s.Push(63)
	s.Push(64)
	s.Push(65)
	s.Push(66)
	s.Push(67)
	s.Push(68)
	s.Push(69)
	s.Push(70)
	s.Push(71)
	s.Push(72)
	s.Push(73)
	s.Push(74)
	s.Push(75)
	s.Push(76)
	s.Push(77)
	s.Push(78)
	s.Push(79)
	s.Push(80)
	s.Push(81)
	s.Push(82)
	s.Push(83)
	s.Push(84)
	s.Push(85)
	s.Push(86)
	s.Push(87)
	s.Push(88)
	s.Push(89)
	s.Push(90)
	s.Push(91)
	s.Push(92)
	s.Push(93)
	s.Push(94)
	s.Push(95)
	s.Push(96)
	s.Push(97)
	s.Push(98)
	s.Push(99)
	for i := 99; i >= 1; i-- {
		x, err := s.Pop()
		if err != nil {
			t.Error("unexpected error")
		}
		if x != i {
			t.Error("expected", i)
		}
	}
}

func BenchmarkMaxStack_Push(b *testing.B) {
	s := MaxStack{}
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}
