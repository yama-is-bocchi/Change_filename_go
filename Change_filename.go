package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
	Dirpath string
}

// フィルター
type Filter func(string) bool

// DirPath,FileNameをフィルターする関数
func FilterName(name string, filter Filter) bool {

	if filter(name) == true {
		fmt.Printf("無効な文字列です\n")
		return false
	}

	return true
}

// ディレクトリ名
func DirFilter() Filter {
	return func(n string) bool {
		if strings.Contains(n, "*") ||
			strings.Contains(n, "?") || strings.Contains(n, "<") ||
			strings.Contains(n, ">") || strings.Contains(n, "\"") {
			return true
		}
		return false
	}
}

// ファイル名
func FileFilter() Filter {
	return func(n string) bool {
		if strings.Contains(n, ".") || strings.Contains(n, "*") ||
			strings.Contains(n, "?") || strings.Contains(n, "<") ||
			strings.Contains(n, ">") || strings.Contains(n, "\"") ||
			strings.Contains(n, "\\") || strings.Contains(n, "/") ||
			strings.Contains(n, "￥") || strings.Contains(n, ":") {
			return true
		}
		return false
	}
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
	for {
		fmt.Printf("入力は間違いないですか?1=OK,0=やり直し\n")
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

// 入力を続けるか確認
// True続ける,False終了
func continue_check() bool {
	fmt.Printf("終了しますか?1=終了,0=続ける\n")
	var ok_str string //入力確認文字列
	for {

		fmt.Scanln(&ok_str)
		if ok_str == "1" {
			return true
		} else if ok_str == "0" {
			return false
		}
		fmt.Printf("入力文字が無効です\n")
	}
}

func main() {
	for {
		// フォルダのパスを指定します
		var FolderList []Path
		var folderPath string    //フォルダのパス
		var newName string       //新規のファイルの名前
		var extensionName string //ファイルのフォーマットの名前

		pathFilter := DirFilter()

		for {
			fmt.Printf("フォルダのパスを入力してください\n")
			folderPath = enter_str()
			if len(folderPath) > 0 && FilterName(folderPath, pathFilter) == true {
				temp := Path{Dirpath: folderPath}
				FolderList = append(FolderList, temp)
				//複数フォルダーを選択するか
				if continue_check() == true {
					break
				}
			}
		}

		fileFilter := FileFilter()

		// 新しい名前に使用するテンプレートを指定します
		for {
			fmt.Printf("ファイルの名前を入力してください\n")
			newName = enter_str()
			if len(newName) > 0 && FilterName(newName, fileFilter) == true {
				break
			}
		}

		// 新しい拡張子を指定します
		for {
			fmt.Printf("拡張子の名前を入力してください(.は含めないでください)\n")
			extensionName = enter_str()
			if len(extensionName) > 0 && FilterName(extensionName, fileFilter) == true {
				break
			}
		}

		//新規のファイルの名前
		NAME := newName + "(%d)." + extensionName

		// 各ファイルに対して名前を変更します
		for _, dir := range FolderList {
			var count int
			count = 1
			// フォルダ内のファイル一覧を取得します
			files, err := filepath.Glob(filepath.Join(dir.Dirpath, "*"))
			if err != nil {
				fmt.Println("ファイルが存在しません", err)
				break
			}
			for _, file := range files {

				info, err := os.Stat(file)
				if err != nil {
					fmt.Println("エラー:", err)
					continue
				}
				if !info.IsDir() {
					// 新しい名前を生成します
					newFileName := fmt.Sprintf(NAME, count)
					count++
					// ファイル名を変更します
					err := os.Rename(file, filepath.Join(dir.Dirpath, newFileName))
					if err != nil {
						fmt.Printf("ファイル名の変更に失敗しました%s: %v\n", file, err)
						continue
					}

					fmt.Printf("Renamed %s to %s\n", file, newFileName)
				}
			}
		}

		if continue_check() == true {
			break
		}
	}
}
