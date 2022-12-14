package main

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFizzbuzz(t *testing.T) {
	Convey("Test Fizzbuzz", t, func() {
		Convey("Write Literal", func() {
			table := []struct {
				n int
				s string
			}{
				{n: 1, s: "1"},
				{n: 9, s: "9"},
				{n: 10, s: "10"},
				{n: 24, s: "24"},
				{n: 245, s: "245"},
				{n: 1004500746, s: "1004500746"},
			}

			literalBuffer := [maxIntLen]byte{}

			for _, c := range table {
				var buffer bytes.Buffer
				writeLiteral(0, c.n, &buffer, literalBuffer[:])

				So(buffer.String(), ShouldEqual, c.s)
			}
		})

		table := []struct {
			start    int
			end      int
			expected string
		}{
			{
				start:    1,
				end:      1,
				expected: "1",
			},
			{
				start:    1,
				end:      2,
				expected: "1, 2",
			},
			{
				start:    1,
				end:      3,
				expected: "1, 2, Fizz",
			},
			{
				start:    1,
				end:      5,
				expected: "1, 2, Fizz, 4, Buzz",
			},
			{
				start:    8,
				end:      14,
				expected: "8, Fizz, Buzz, 11, Fizz, 13, 14",
			},
			{
				start:    8,
				end:      20,
				expected: "8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16, 17, Fizz, 19, Buzz",
			},
		}

		for _, c := range table {
			var buffer bytes.Buffer
			err := FizzBuzz(c.start, c.end, []byte(", "), &buffer)
			So(err, ShouldBeNil)
			SoMsg(fmt.Sprintf(`start = %d, end = %d, expected = "%s"`, c.start, c.end, c.expected), buffer.String(), ShouldEqual, c.expected)
		}
	})
}
