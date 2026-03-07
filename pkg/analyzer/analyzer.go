package analyzer

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/OWEEN3/loglint/pkg/analyzer/rules"
	"golang.org/x/tools/go/analysis"
)

const AnalyzerName = "loglint"

var Analyzer = &analysis.Analyzer{
	Name: AnalyzerName,
	Doc:  "checks logging format consistency across the project",
	Run:  run,
}

// Запускаем анализатор
func run(pass *analysis.Pass) (interface{}, error) {
	// Методы которые мы будем проверять у slog
	slogMethods := map[string]struct{}{
		"Info": {}, "Error": {}, "Debug": {}, "Warn": {},
	}

	// Методы которые мы будем проверять у zap
	zapMethods := map[string]struct{}{
		"Info": {},
		"Error": {}, 
		"Debug": {}, 
		"Warn": {}, 
		"Panic": {}, 
		"Fatal": {}, 

		// Их нужно обрабатывать отдельно, в случае с fками ругается на спец символы
		// "Infof": {}, "Infow": {},
		// "Errorf": {}, "Errorw": {},
		// "Debugf": {}, "Debugw": {},
		// "Warnf": {}, "Warnw": {},
		// "Panicf": {}, "Panicw": {},
		// "Fatalf": {}, "Fatalw": {},
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// проверяем только вызовы функций
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			// 
			var obj types.Object
			switch fun := call.Fun.(type) {
			// Если импортировали slog например так import . "log/slog" и далее вызвали Info()
			case *ast.Ident:
				obj = pass.TypesInfo.ObjectOf(fun)
			case *ast.SelectorExpr:
				obj = pass.TypesInfo.ObjectOf(fun.Sel)
			default:
				return true
			}

			if obj == nil {
				return true
			}

			pkgPath := ""
			if obj.Pkg() != nil {
				pkgPath = obj.Pkg().Path()
			}

			switch pkgPath {
			case "log/slog":
				if _, ok := slogMethods[obj.Name()]; ok {
					checkLogMessage(pass, call)
				}
			case "go.uber.org/zap":
				if _, ok := zapMethods[obj.Name()]; ok {
					checkLogMessage(pass, call)
				}
			}

			return true
		})
	}

	return nil, nil
}

func checkLogMessage(pass *analysis.Pass, call *ast.CallExpr) {
	// Пустой вызов проверяем, на всякий случай
	if len(call.Args) == 0 {
		pass.Reportf(call.Pos(), "call without arguments")
		return
	}

	// Проходимся по аргументам
	for _, arg := range call.Args {
		checkArg(pass, arg)
	}
}

func checkArg(pass *analysis.Pass, expr ast.Expr) {
	switch e := expr.(type) {
	// Если обычная строка
	case *ast.BasicLit:
		msg := strings.Trim(e.Value, "\"`")
		if rules.IsEmpty(msg) {
			pass.Reportf(e.Pos(), "empty message %s", msg)
			return
		}
		if !rules.IsValidChars(msg) {
			pass.Reportf(e.Pos(), "invalid letters %s", msg)
			return
		}
		if !rules.IsFirstLower(msg) {
			pass.Reportf(e.Pos(), "first letter should be lowercase %s", msg)
			return
		}
	// Если была конкатенация
	case *ast.BinaryExpr:
		checkArg(pass, e.X)
		checkArg(pass, e.Y)
	// Проверяем название переменной на наличие чувствительных данных
	case *ast.Ident:
		if str := rules.ContainsSensitive(e.Name); str != "" {
			pass.Reportf(e.Pos(), "sensitive data may be stored %s; %s", str, e.Name)
		}
	// Рекурсивно вызываем для случаев типо zap.String("key", "value") и т.д.
	case *ast.CallExpr:
		checkLogMessage(pass, e)
	default:
		pass.Reportf(expr.Pos(), "unsupported argument type")
	}
}