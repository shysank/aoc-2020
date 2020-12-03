package puzzles

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Puzzles", func() {

	const (
		inputFilesBaseDir = "../inputs"
	)

	Context("Day 1", func() {
		subject := &reportRepair{2020}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day1/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("1010884"))

		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day1/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("253928438"))
		})
	})

	Context("Day 2", func() {
		subject := &passwordPhilosophy{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day2/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("580"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day2/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("611"))
		})
	})

	Context("Day 3", func() {
		subject := &tobogganTrajectory{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day3/p1")
			Expect(err).To(BeNil())

			subject.slopes = []tobogganCoordinates{{3, 1}}
			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("223"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day3/p1")
			Expect(err).To(BeNil())

			subject.slopes = []tobogganCoordinates{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("3517401300"))
		})
	})

})
