package internal

import (
	"io/ioutil"
	"path"
)

func ScanEnv(root string) ([]string, error) {

	files, err := ioutil.ReadDir(root)

	if err != nil {
		return nil, err
	}

	env := make([]string, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, _ := ioutil.ReadFile(path.Join(root, file.Name()))
		arg := file.Name() + "=" + string(content)
		env = append(env, arg)
	}

	return env, nil
}
