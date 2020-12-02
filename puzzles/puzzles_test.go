package puzzles

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Puzzles", func() {

	const (
		inputFilesBaseDir = "../inputs"
	)

	var (
		subject Solver
	)
	Context("Day 1", func() {
		subject = &expenseReport{2020}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day1/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())

			p1Result := result.day1.p1
			Expect(p1Result.n1 + p1Result.n2).Should(Equal(int64(2020)))
			fmt.Println(p1Result.n1 * p1Result.n2)

		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day1/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())

			p2Result := result.day1.p2
			Expect(p2Result.n1 + p2Result.n2 + p2Result.n3).Should(Equal(int64(2020)))
			fmt.Println(p2Result.n1 * p2Result.n2 * p2Result.n3)

		})

	})
})
