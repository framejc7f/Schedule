package WorkingWithFiles

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/framejc7f/Schedule/structs"
	excelize "github.com/xuri/excelize/v2"
)

func ReadFile(fileName string) []structs.Week {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheets := f.GetSheetList()
	lists := []structs.Week{}
	for i, v := range sheets {
		lists = append(lists, structs.Week{Id: strconv.Itoa(i), Name: v})
	}

	return lists

	// fmt.Println("Выбор недели:")
	// for i, v := range sheets {
	// 	fmt.Printf("%d) %s\n", i+1, v)
	// }
	// var inp string
	// fmt.Print("--> ")
	// fmt.Scanf("%s\n", &inp)
	// i, _ := strconv.Atoi(inp)
	// week := sheets[i-1]
	// fmt.Printf("Выбор --- %s\n", week)

	// cell, err := f.GetCellValue(, "B2")
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Println(cell)
}

func cell(f *excelize.File, sheet, val string) string {
	cell, err := f.GetCellValue(sheet, val)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return cell
}

func ReadDay(fileName, sheet string, day int, subgroup int) structs.Day {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	data := cell(f, sheet, fmt.Sprintf("AI%d", 7+8*day))
	var lessons []structs.Lesson
	for i := 1; i <= 8; i++ {
		time := cell(f, sheet, fmt.Sprintf("AK%d", 7+8*day+i-1))
		var subject string
		var typ string
		var classroom string
		if subgroup == 1 {
			subject = cell(f, sheet, fmt.Sprintf("AL%d", 7+8*day+i-1))
			typ = cell(f, sheet, fmt.Sprintf("AM%d", 7+8*day+i-1))
			classroom = cell(f, sheet, fmt.Sprintf("AN%d", 7+8*day+i-1))
		} else if subgroup == 2 {
			subject = cell(f, sheet, fmt.Sprintf("AO%d", 7+8*day+i-1))
			typ = cell(f, sheet, fmt.Sprintf("AP%d", 7+8*day+i-1))
			classroom = cell(f, sheet, fmt.Sprintf("AQ%d", 7+8*day+i-1))
		} else {
			subject = cell(f, sheet, fmt.Sprintf("AR%d", 7+8*day+i-1))
			typ = cell(f, sheet, fmt.Sprintf("AS%d", 7+8*day+i-1))
			classroom = cell(f, sheet, fmt.Sprintf("AT%d", 7+8*day+i-1))
		}
		if typ == classroom && typ != "" && classroom != "" {
			typ = cell(f, sheet, fmt.Sprintf("AS%d", 7+8*day+i-1))
			classroom = cell(f, sheet, fmt.Sprintf("AT%d", 7+8*day+i-1))
		}

		if subject == "" {
			continue
		} else {
			subject = strings.TrimSpace(subject)
			typ = strings.TrimSpace(typ)
			classroom = strings.TrimSpace(classroom)
			lesson := structs.Lesson{Id: strconv.Itoa(i), Time: time, Subgroup: subgroup, Subject: subject, Сlassroom: classroom, Type: typ}
			lessons = append(lessons, lesson)
		}
	}
	Day := structs.Day{Id: strconv.Itoa(day), Data: data, Lessons: lessons}
	return Day
}
