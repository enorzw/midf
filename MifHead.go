/*
MIF文件头
AddColumn(cname string, ctype string) bool
String() string
*/

package midf

import (
	"fmt"
	"strings"
)

const (
	STR_NEWLINE = "\n"
	STR_SPACE   = " "
	STR_QUOTE   = "\""
	STR_COMMA   = ","
	STR_TAB     = "\t"
)

// char(width)
// integer (which is 4 bytes)
// smallint (which is 2 bytes, so it can only store numbers between
//   -32767 and +32767)
// decimal (width, decimals)
// float
// date
// logical
const (
	COLUMNTYPE_CHAR = iota
	COLUMNTYPE_INTEGER
	COLUMNTYPE_SMALLINT
	COLUMNTYPE_DECIMAL
	COLUMNTYPE_FLOAT
	COLUMNTYPE_DATE
	COLUMNTYPE_LOGICAL
)

type MifHead struct {
	Version   int
	Charset   string
	Delimiter string
	Unique    []int
	Index     []int
	Coordsys  string
	Transform string

	columnNumber  int
	columnNames   []string
	columnType    []string
	columnNameMap map[string]int
}

func NewMifHead() MifHead {
	head := MifHead{}
	head.Version = 300
	head.Charset = "WindowsSimpChinese"
	head.Delimiter = "\t"
	head.columnNumber = 0
	head.columnNames = make([]string, 0, 10)
	head.columnType = make([]string, 0, 10)
	head.columnNameMap = make(map[string]int)
	return head
}

func (m *MifHead) AddColumn(cname string, ctype string) bool {
	if _, ok := m.columnNameMap[cname]; ok {
		return false
	} else {
		m.columnNameMap[cname] = m.columnNumber
		m.columnNames = append(m.columnNames, cname)
		m.columnType = append(m.columnType, ctype)
		m.columnNumber += 1
	}
	return true
}

func (m *MifHead) String() string {
	buffer := ""
	//Version
	version := fmt.Sprintf("%d", m.Version)
	buffer += "Version" + STR_SPACE + version + STR_NEWLINE
	//Charset
	buffer += "Charset" + STR_SPACE + STR_QUOTE + m.Charset + STR_QUOTE + STR_NEWLINE
	//Delimiter
	buffer += "Delimiter" + STR_SPACE + STR_QUOTE + m.Delimiter + STR_QUOTE + STR_NEWLINE
	//Unique
	if m.Unique != nil && len(m.Unique) > 0 {
		buffer += "Unique" + STR_SPACE
		for i := 0; i < len(m.Unique); i++ {
			str := fmt.Sprintf("%d", m.Unique[i])
			buffer += str + STR_COMMA
		}
		buffer = strings.TrimRight(buffer, STR_COMMA)
		buffer += STR_NEWLINE
	}
	//Index
	if m.Index != nil && len(m.Index) > 0 {
		buffer += "Index" + STR_SPACE
		for i := 0; i < len(m.Index); i++ {
			str := fmt.Sprintf("%d", m.Index[i])
			buffer += str + STR_COMMA
		}
		buffer = strings.TrimRight(buffer, STR_COMMA)
		buffer += STR_NEWLINE
	}
	//Coordsys
	if m.Coordsys != "" {
		buffer += "Coordsys" + STR_SPACE + m.Coordsys + STR_NEWLINE
	}
	//Transform
	if m.Transform != "" {
		buffer += "Transform" + STR_SPACE + m.Transform + STR_NEWLINE
	}
	//Columns
	columns := fmt.Sprintf("%d", m.columnNumber)
	buffer += "Columns" + STR_SPACE + columns + STR_NEWLINE
	for i := 0; i < m.columnNumber; i++ {
		buffer += STR_TAB + m.columnNames[i] + STR_SPACE + m.columnType[i] + STR_NEWLINE
	}
	//Data
	buffer += "Data" + STR_NEWLINE
	return buffer
}
