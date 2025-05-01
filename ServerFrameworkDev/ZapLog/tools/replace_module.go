package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// replaceImportsInFile 函数用于读取指定文件内容，将其中的旧模块名替换为新模块名，然后将修改后的内容写回文件。
// 参数 filePath 为要处理的文件的路径，oldModule 为原模块名，newModule 为要替换成的模块名。
// 返回值为处理过程中可能出现的错误，如果处理成功则返回 nil。
func replaceImportsInFile(filePath string, oldModule, newModule string) error {
	// 读取指定文件的全部内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 将文件内容中的旧模块名替换为新模块名
	newContent := strings.ReplaceAll(string(content), oldModule, newModule)
	// 将替换后的内容写回原文件，文件权限设置为 0644
	return os.WriteFile(filePath, []byte(newContent), 0644)
}

// walkDir 函数用于递归遍历指定根目录下的所有文件和子目录。
// 对于其中的 .go 文件，调用 replaceImportsInFile 函数替换模块导入名。
// 参数 root 为要遍历的根目录路径，oldModule 为原模块名，newModule 为要替换成的模块名。
// 返回值为遍历过程中可能出现的错误，如果遍历成功则返回 nil。
func walkDir(root, oldModule, newModule string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			err := replaceImportsInFile(path, oldModule, newModule)
			if err != nil {
				return err
			}
			fmt.Printf("已更新来自于 %s 的导入项\n", path)
		}
		return nil
	})
}

// replaceModuleInGoMod 函数用于替换 go.mod 文件中的 module 名
// 参数 goModPath 为 go.mod 文件的路径，oldModule 为原模块名，newModule 为要替换成的模块名。
// 返回值为处理过程中可能出现的错误，如果处理成功则返回 nil。
func replaceModuleInGoMod(goModPath, oldModule, newModule string) error {
	// 读取 go.mod 文件的全部内容
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return err
	}

	// 构造 module 声明的旧字符串和新字符串
	oldModuleLine := fmt.Sprintf("module %s", oldModule)
	newModuleLine := fmt.Sprintf("module %s", newModule)

	// 将 go.mod 文件内容中的旧 module 声明替换为新 module 声明
	newContent := strings.ReplaceAll(string(content), oldModuleLine, newModuleLine)
	// 将替换后的内容写回 go.mod 文件，文件权限设置为 0644
	return os.WriteFile(goModPath, []byte(newContent), 0644)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 提示用户输入原模块名
	fmt.Print("请输入原模块名: ")
	scanner.Scan()
	oldModule := scanner.Text()

	// 提示用户输入要替换成的模块名
	fmt.Print("请输入要替换成的模块名: ")
	scanner.Scan()
	newModule := scanner.Text()

	// 提示用户输入要遍历的目录路径，添加示例和说明（\"：表示双引号。 \\：表示反斜杠。）
	fmt.Print("请确定要遍历的目录路径，\n如 E:\\Folder01\\Folder02\\Folder03\\Project，若有路径空格，请将整个路径用英文格式的双引号 \"\" 框起来。\n请输入路径: ")
	scanner.Scan()
	rootDir := scanner.Text()

	// 去除用户输入可能包含的双引号
	rootDir = strings.Trim(rootDir, "\"")

	// 构造 go.mod 文件的路径
	goModPath := filepath.Join(rootDir, "go.mod")

	// 确认用户操作
	fmt.Printf("即将在目录 %s 下的所有 .go 文件以及 go.mod 文件中，将 %s 替换为 %s。\n确认操作请输入 y，取消请输入其他任意键: ", rootDir, oldModule, newModule)
	scanner.Scan()
	confirm := scanner.Text()
	if strings.ToLower(confirm) != "y" {
		fmt.Println("操作已取消。")
		return
	}

	// 替换 go.mod 文件中的 module 名
	err := replaceModuleInGoMod(goModPath, oldModule, newModule)
	if err != nil {
		fmt.Printf("更新 go.mod 文件时出错: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("go.mod 中 module字段 已更新。")

	// 调用 walkDir 函数递归遍历指定目录下的所有 .go 文件，替换其中的模块导入名
	err = walkDir(rootDir, oldModule, newModule)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("所有导入项已成功更新。")
}
