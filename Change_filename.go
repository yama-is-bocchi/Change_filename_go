package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//入力文字列に無効な文字が入力されていないか確認
//戻り値True:OK,False:NOT
func check_str(input string)bool{

	if !strings.Contains(input,"￥") || !strings.Contains(input,"/")||
	!strings.Contains(input,".")|| !strings.Contains(input,"\\")||
	!strings.Contains(input,":")|| !strings.Contains(input,"\"")||
	!strings.Contains(input,"*")|| !strings.Contains(input,"?")||
	!strings.Contains(input,"<")|| !strings.Contains(input,">"){
		return true
	}
	fmt.Printf("無効な文字列です\n")
	return false
}

// 文字列を入力して返す
// 戻り値input or ""(空文字)
func enter_str() string {

	var input string  //入力文字列
	var ok_str string //入力確認文字列

	for {
		// Enterが押されるまで入力を待つ
		fmt.Scanln(&input)
		if len(input) > 0 {
			break
		}

		fmt.Printf("文字列の長さが0です\n")
	}

	fmt.Println("入力は間違いないですか?1=OK,0=やり直し\n", input)

	for {

		fmt.Scanln(&ok_str)
		if ok_str == "1" {
			return input
		} else if ok_str == "0" {
			break
		}
		fmt.Printf("入力文字が無効です\n")
	}
	return ""
}

func main() {
	// フォルダのパスを指定します
	var folderPath string    //フォルダのパス
	var newName string       //新規のファイルの名前
	var extensionName string //ファイルのフォーマットの名前

	for {
		fmt.Printf("フォルダのパスを入力してください\n")
		folderPath = enter_str()
		if len(folderPath) > 0 {
			break
		}
	}

	// フォルダ内のファイル一覧を取得します
	files, err := filepath.Glob(filepath.Join(folderPath, "*"))
	if err != nil {
		fmt.Println("ファイルが存在しません", err)
		return
	}

	// 新しい名前に使用するテンプレートを指定します
	for {
		fmt.Printf("ファイルの名前を入力してください\n")
		newName = enter_str()
		if len(newName) > 0 ||check_str(newName)==true{
			break
		}
	}

	// 新しい拡張子を指定します
	for {
		fmt.Printf("拡張子の名前を入力してください(.は含めないでください)\n")
		extensionName = enter_str()
		if len(extensionName) > 0 ||check_str(extensionName)==true{
			break
		}
	}

	//新規のファイルの名前
	NAME:=newName+"(%d)."+extensionName

	// 各ファイルに対して名前を変更します
	for i, file := range files {
		// 新しい名前を生成します
		newFileName := fmt.Sprintf(NAME, i+1)

		// ファイル名を変更します
		err := os.Rename(file, filepath.Join(folderPath, newFileName))
		if err != nil {
			fmt.Printf("ファイル名の変更に失敗しました%s: %v\n", file, err)
			continue
		}

		fmt.Printf("Renamed %s to %s\n", file, newFileName)
	}
}
