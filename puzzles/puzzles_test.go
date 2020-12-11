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

			subject.slopes = []coordinates{{3, 1}}
			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("223"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day3/p1")
			Expect(err).To(BeNil())

			subject.slopes = []coordinates{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("3517401300"))
		})
	})

	Context("Day 4", func() {
		subject := NewPassportProcessing()
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day4/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("230"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day4/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("156"))
		})
	})

	Context("Day 5", func() {
		subject := &binaryBoarding{128, 8}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day5/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("871"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day5/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("640"))
		})
	})

	Context("Day 6", func() {
		subject := &customCustoms{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day6/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("6161"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day6/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("2971"))
		})
	})

	Context("Day 7", func() {
		subject := &handyHaversacks{colorToCheck: "shiny gold"}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day7/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("259"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day7/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("45018"))
		})
	})

	Context("Day 8", func() {
		subject := &handheldHalting{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day8/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("1930"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day8/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("1688"))
		})
	})

	Context("Day 9", func() {
		subject := &encodingError{preambleLength: 25}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day9/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("248131121"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day9/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("31580383"))
		})
	})

	Context("Day 10", func() {
		subject := &adapterArray{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day10/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("2376"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day10/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("129586085429248"))
		})
	})

	Context("Day 11", func() {
		subject := &seatingSystem{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day11/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("2481"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day11/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("2227"))
		})
	})

})
