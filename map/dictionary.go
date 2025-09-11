package main

// var (
// 	ErrNotFound   = errors.New("could not find the word you were looking for")
// 	ErrWordExists = errors.New("cannot add the word because it already exists")
// )

// 错误声明为常量
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// 实现了 Error 接口
func (dicErr DictionaryErr) Error() string {
	return string(dicErr)
}

type Dictionary map[string]string

func (dic Dictionary) Search(word string) (string, error) {
	target, ok := dic[word]

	if !ok {
		return "", ErrNotFound
	}

	return target, nil
}

// map 是引用类型，不需要传递指针就可以修改它的值
func (dic Dictionary) Add(key string, val string) error {

	// if _, err := dic.Search(key); err == nil {
	// 	return ErrWordExists
	// } else {
	// 	dic[key] = val
	// 	return nil
	// }

	_, err := dic.Search(key)
	switch err {
	case ErrNotFound:
		dic[key] = val
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (dic Dictionary) Update(word, definition string) error {
	_, err := dic.Search(word)
	switch err {
	case nil:
		dic[word] = definition
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (dic Dictionary) Delete(word string) {
	delete(dic, word)
}
