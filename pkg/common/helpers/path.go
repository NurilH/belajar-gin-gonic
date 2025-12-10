package helpers

import "os"

func GetUploadDir() string {
	if dir := os.Getenv("UPLOAD_DIR"); dir != "" {
		return dir
	}

	if _, err := os.Stat("/app/static/file"); err != nil {
		return "/app/static/file"
	}

	return "./static/file"
}

func GetStaticDir() string {
	if dir := os.Getenv("STATIC_DIR"); dir != "" {
		return dir
	}

	if _, err := os.Stat("/app/static"); err != nil {
		return "/app/static"
	}

	return "./static"
}
