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
		subject := &ReportRepair{2020}
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
		subject := &PasswordPhilosophy{}
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
		subject := &TobogganTrajectory{}
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
		subject := &BinaryBoarding{128, 8}
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
		subject := &CustomCustoms{}
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
		subject := &HandyHaversacks{colorToCheck: "shiny gold"}
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
		subject := &HandheldHalting{}
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
		subject := &EncodingError{preambleLength: 25}
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
		subject := &AdapterArray{}
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
		subject := &SeatingSystem{}
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

	Context("Day 12", func() {
		subject := &RainRisk{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day12/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("1294"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day12/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("20592"))
		})
	})

	Context("Day 13", func() {
		subject := &ShuttleSearch{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day13/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("410"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day13/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("600691418730595"))
		})
	})

	Context("Day 14", func() {
		subject := &DockingData{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day14/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("17481577045893"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day14/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("4160009892257"))
		})
	})

	Context("Day 15", func() {
		subject := &RambunctiousRecitation{[]int{5, 1, 9, 18, 13, 8, 0}, 2020}
		It("puzzle 1", func() {
			result, err := subject.Puzzle1(nil)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("376"))
		})

		It("puzzle 2", func() {
			subject := &RambunctiousRecitation{[]int{5, 1, 9, 18, 13, 8, 0}, 30000000}
			result, err := subject.Puzzle2(nil)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("323780"))
		})
	})

	Context("Day 16", func() {
		subject := &TicketTranslation{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day16/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("30869"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day16/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("4381476149273"))
		})

	})

	Context("Day 17", func() {
		subject := &ConwayCubes{cycles: 6}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day17/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("276"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day17/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("2136"))
		})

	})

	Context("Day 18", func() {
		subject := &OperationOrder{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day18/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("45283905029161"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day18/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("216975281211165"))
		})

	})

	Context("Day 19", func() {
		subject := &MonsterMessages{}
		It("puzzle 1", func() {
			file, err := os.Open(inputFilesBaseDir + "/day19/p1")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle1(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("272"))
		})

		It("puzzle 2", func() {
			file, err := os.Open(inputFilesBaseDir + "/day19/p2")
			Expect(err).To(BeNil())

			result, err := subject.Puzzle2(file)
			Expect(err).To(BeNil())
			Expect(result.Value()).Should(Equal("374"))
		})

	})
})
