package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CronParser is the interface for parsing cron strings.
type CronParser interface {
	Parse(cronString string) (CronFields, error)
}

// CronFields stores the parsed fields of a cron string.
type CronFields struct {
	Minute     []string
	Hour       []string
	DayOfMonth []string
	Month      []string
	DayOfWeek  []string
	Command    string
}

// Parser is the struct implementing the CronParser interface.
type Parser struct{}

// NewParser returns a new instance of Parser.
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses a cron string into its fields.
func (p *Parser) Parse(cronString string) (CronFields, error) {
	parts := strings.Fields(cronString)
	if len(parts) != 6 {
		return CronFields{}, fmt.Errorf("invalid cron string format. Expected 5 fields and a command")
	}

	minute := parts[0]
	hour := parts[1]
	dayOfMonth := parts[2]
	month := parts[3]
	dayOfWeek := parts[4]
	command := parts[5]

	minutes := expandField(minute, "0", "59")
	hours := expandField(hour, "0", "23")
	daysOfMonth := expandField(dayOfMonth, "1", "31")
	months := expandField(month, "1", "12")
	daysOfWeek := expandField(dayOfWeek, "0", "6")

	return CronFields{
		Minute:     minutes,
		Hour:       hours,
		DayOfMonth: daysOfMonth,
		Month:      months,
		DayOfWeek:  daysOfWeek,
		Command:    command,
	}, nil
}

func expandField(field, min, max string) []string {
	var result []string
	if field == "*" {
		for i := toInt(min); i <= toInt(max); i++ {
			result = append(result, strconv.Itoa(i))
		}
	} else if strings.Contains(field, ",") {
		for _, part := range strings.Split(field, ",") {
			result = append(result, expandField(part, min, max)...)
		}
	} else if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		start := toInt(parts[0])
		end := toInt(parts[1])
		for i := start; i <= end; i++ {
			result = append(result, strconv.Itoa(i))
		}
	} else if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		base := parts[0]
		step := toInt(parts[1])
		if base == "*" {
			for i := toInt(min); i <= toInt(max); i += step {
				result = append(result, strconv.Itoa(i))
			}
		} else {
			start := toInt(base)
			for i := start; i <= toInt(max); i += step {
				result = append(result, strconv.Itoa(i))
			}
		}
	} else {
		result = append(result, field)
	}
	return result
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cron_parser \"<cron_string>\"")
		return
	}

	cronString := os.Args[1]
	parser := NewParser()
	parsedFields, err := parser.Parse(cronString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("minute        %s\n", strings.Join(parsedFields.Minute, " "))
	fmt.Printf("hour          %s\n", strings.Join(parsedFields.Hour, " "))
	fmt.Printf("day of month  %s\n", strings.Join(parsedFields.DayOfMonth, " "))
	fmt.Printf("month         %s\n", strings.Join(parsedFields.Month, " "))
	fmt.Printf("day of week   %s\n", strings.Join(parsedFields.DayOfWeek, " "))
	fmt.Printf("command       %s\n", parsedFields.Command)
}
