/**
 * @Author pibing
 * @create 2020/11/15 10:37 AM
 */

package excel



import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
    "github.com/plandem/xlsx"
)


/**
 * 导出Excel表格
 * @param  name     {string}    导出的表名
 * @param  header   {[]string}  表头key，导出后显示的顺序
 * @param  headerKV {map[string]string}  表头、数据kv对照
 * @param  data     {[]map[string]interface{}} 数据集合
 * @return err      {error}                    异常
 */
func ExportExcel(name string, header []string, headerKV map[string]string, data []map[string]interface{}) (fileName string, err error) {
	f := excelize.NewFile()
	// Create a new sheet
	index := f.NewSheet("Sheet1")
	headers := make([]string, 0)
	for _, v := range header {
		headers = append(headers, headerKV[v])
	}
	f.SetSheetRow("Sheet1", "A1", &headers)
	var rowValue []interface{}
	for i, v := range data {
		rowValue = make([]interface{}, 0)
		//表中的行顺序是从1开始的;  A1是表头,A2才是数据的开始行
		rowNum := strconv.Itoa(i + 2)
		for _, key := range header {
			rowValue = append(rowValue, v[key])
		}
		f.SetSheetRow("Sheet1", "A"+rowNum, &rowValue)
	}

	f.SetActiveSheet(index)
	fileName = name + ".xlsx"       //文件名称
	fileNamePath := "./" + fileName  //保存文件的位置
	err = f.SaveAs(fileNamePath)
	return
}



/*
读取excel表中数据,返回list，可改用结构体list
*/
func UploadExcel(filePath string,head []string)([]map[string]interface{}) {
	list := make([]map[string]interface{},0)
	if len(head)==0 {
		return list
	}
	xl, err := xlsx.Open(filePath)
	if err != nil {
		fmt.Print("打开文件err：",err)
		return list
	}
	defer xl.Close()
	sheet := xl.Sheet(0, xlsx.SheetModeIgnoreDimension)  //读取模式
	_, totalRows := sheet.Dimension()
	for row := 1; row < totalRows; row++ {
		rowmap:=make(map[string]interface{})
		for col := 0; col < len(head); col++ {
			value := sheet.Cell(col, row).String()
			field :=head[col]
			rowmap[field]=value
		}
		list=append(list,rowmap)
	}
	return list
}
