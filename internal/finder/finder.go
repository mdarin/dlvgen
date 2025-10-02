package finder

// Установка: go get github.com/jaytaylor/go-find
import (
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/jaytaylor/go-find"
)

// FindMainProgram ищет точку входа main.go в проекте.
// Он сначала собирает всех кандидатов, затем приоритизирует их.
func FindMainProgram(paths []string) string {
	slog.Debug("Searching for main.go files with priority")

	// try to find from workspace root (.)
	finder := find.NewFind(paths...).MinDepth(0).Name("main.go")
	hits, err := finder.Evaluate()
	if err != nil {
		panic(err)
	}

	if len(hits) == 0 {
		slog.Warn("No main.go files found, using default")
		return "./main.go"
	}

	slog.Debug("Found candidates", "files", hits)

	// Find the Least Nested main.go File
	// slices.SortFunc(hits, comparePathDepths)

	// 3. Логика приоритизации

	// 3a. Сначала ищем приоритетный путь в /cmd/
	cmdSeparator := "cmd" + string(filepath.Separator)
	for _, candidate := range hits {
		if strings.Contains(candidate, cmdSeparator) {
			slog.Info("Selected main program from high-priority path", "file", candidate)
			return candidate
		}
	}

	// 3b. Если в /cmd/ ничего нет, сортируем всех по глубине, как вы предложили.
	// Используем slices.SortFunc для чистоты кода.
	slices.SortFunc(hits, func(a, b string) int {
		depthA := strings.Count(a, string(filepath.Separator))
		depthB := strings.Count(b, string(filepath.Separator))
		// Сравниваем для сортировки по возрастанию глубины
		if depthA < depthB {
			return -1
		}
		if depthA > depthB {
			return 1
		}
		return 0
	})

	// После сортировки лучший кандидат (самый неглубокий) будет первым.
	bestCandidate := hits[0]
	slog.Info("Selected main program by shortest path", "file", bestCandidate)

	return bestCandidate
}

// func findMainProgram() string {
// 	slog.Debug("Searching for main.go files using go-find")

// 	// Ищем файлы main.go, содержащие "package main"
// 	finder, err := find.New(
// 		".", // Искать в текущей директории и поддиректориях
// 		find.WithRegexps("main.go$"),
// 		find.WithContaining("package main", "func main()"),
// 	)
// 	if err != nil {
// 		slog.Error("Failed to initialize file finder", "error", err)
// 		return "./main.go" // Возвращаем безопасное значение по умолчанию
// 	}

// 	var candidates []string
// 	// Результаты приходят асинхронно
// 	for res := range finder.Find() {
// 		if res.Err != nil {
// 			slog.Debug("Error during file search", "path", res.Path, "error", res.Err)
// 			continue
// 		}
// 		candidates = append(candidates, res.Path)
// 		slog.Debug("Found main program candidate", "file", res.Path)
// 	}

// 	if len(candidates) == 0 {
// 		slog.Warn("No main.go files found, using default")
// 		return "./main.go"
// 	}

// 	// Та же логика приоритизации, что и у вас
// 	bestCandidate := candidates[0]
// 	minDepth := strings.Count(bestCandidate, string(filepath.Separator))

// 	for _, candidate := range candidates {
// 		// Приоритет для путей, содержащих /cmd/
// 		if strings.Contains(candidate, "cmd"+string(filepath.Separator)) {
// 			slog.Info("Selected main program from priority path", "file", candidate)
// 			return candidate
// 		}
// 		depth := strings.Count(candidate, string(filepath.Separator))
// 		if depth < minDepth {
// 			minDepth = depth
// 			bestCandidate = candidate
// 		}
// 	}

// 	slog.Info("Selected main program by shortest path", "file", bestCandidate)
// 	return bestCandidate
// }

func findMainProgramObsolete() string {
	slog.Debug("Searching for main.go files")

	searchPaths := []string{
		".",
		"cmd",
		"app",
		"src",
		"main",
	}

	var candidates []string

	for _, path := range searchPaths {
		err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			if !info.IsDir() && strings.HasSuffix(filePath, ".go") {
				content, err := os.ReadFile(filePath)
				if err == nil {
					if strings.Contains(string(content), "package main") &&
						strings.Contains(string(content), "func main()") {
						candidates = append(candidates, filePath)
						slog.Debug("Found main program candidate", "file", filePath)
					}
				}
			}
			return nil
		})

		if err != nil {
			slog.Debug("Error walking path", "path", path, "error", err)
		}
	}

	if len(candidates) > 0 {
		// Prioritize by depth and common patterns
		bestCandidate := candidates[0]
		for _, candidate := range candidates {
			if strings.Contains(candidate, "cmd/") || strings.Contains(candidate, "main/") {
				bestCandidate = candidate
				break
			}
			if strings.Count(candidate, string(filepath.Separator)) < strings.Count(bestCandidate, string(filepath.Separator)) {
				bestCandidate = candidate
			}
		}
		slog.Info("Selected main program", "file", bestCandidate)
		return bestCandidate
	}

	slog.Warn("No main.go files found, using default")
	return "./main.go"
}
