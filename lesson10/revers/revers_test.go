package revers

import "testing"

var (
	list = OneWayList{
		value: 5,
		next: &OneWayList{
			value: 4,
			next: &OneWayList{
				value: 3,
				next: &OneWayList{
					value: 2,
					next: &OneWayList{
						value: 1,
						next: &OneWayList{
							value: 0,
						},
					},
				},
			},
		},
	}
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func BenchmarkOneWayList_ReversByDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list.ReversByDefer()
	}
}

func BenchmarkOneWayList_ReverseByLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list.ReverseByLoop()
	}
}

func BenchmarkOneWayList_ReversByRecurse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list.ReversByRecurse()
	}
}
