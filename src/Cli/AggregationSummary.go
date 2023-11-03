package Cli

import (
    "strconv"
    "fmt"
    "github.com/halleck45/ast-metrics/src/Analyzer"
    "github.com/charmbracelet/glamour"
)

func AggregationSummary(aggregated Analyzer.Aggregated) {

    style := StyleTitle()
    fmt.Println(style.Render("Results overview"))

   in := `*This code is composed from ` + strconv.Itoa(aggregated.Loc) + ` lines of code, ` + strconv.Itoa(aggregated.Cloc) + ` comment lines of code and ` + strconv.Itoa(aggregated.Lloc) + ` logical lines of code.*

   ## Complexity

   ### Cyclomatic complexity

   *Cyclomatic Complexity is a measure of the number of linearly independent paths through a program's source code.
   More you have paths, more your code is complex.*


   | Min | Max | Average per class | Average per method |
   | --- | --- | --- | --- |
   | ` +
        strconv.Itoa(aggregated.MinCyclomaticComplexity) +
        ` | ` + strconv.Itoa(aggregated.MaxCyclomaticComplexity) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.AverageCyclomaticComplexityPerClass) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.AverageCyclomaticComplexityPerMethod) +
        ` |


   ### Halstead metrics

   *Halstead metrics are software metrics introduced to empirically determine the complexity of a program.*

   | | Difficulty | Effort | Volume | Time |
   | --- | --- | --- | --- | --- |
    ` +
        ` | Total` +
        ` | ` + fmt.Sprintf("%.2f", aggregated.SumHalsteadDifficulty) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.SumHalsteadEffort) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.SumHalsteadVolume) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.SumHalsteadTime) +
        "\n | Average per class" +
        ` | ` + fmt.Sprintf("%.2f", aggregated.AverageHalsteadDifficulty) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.AverageHalsteadEffort) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.AverageHalsteadVolume) +
        ` | ` + fmt.Sprintf("%.2f", aggregated.AverageHalsteadTime) +
        ` |

   ### Classes and methods

   | Classes | Methods | Average methods per class |
   | --- | --- | --- |
   | ` + strconv.Itoa(aggregated.NbClasses) + ` | ` + strconv.Itoa(aggregated.NbMethods) + ` | ` + fmt.Sprintf("%.2f", aggregated.AverageMethodsPerClass) + ` |

   `
   out, _ := glamour.Render(in, "dark")
   fmt.Print(out)
}