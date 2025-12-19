package documentstore

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type DocumentFieldType
	// редактор пораздив замінити interface{} на any
	Value any
}

type Document struct {
	Fields map[string]DocumentField
}

var documents = map[string]*Document{}

// функція Put для додавання і оновлення данних у мапі, поле key обовʼязкове

func Put(doc *Document) {
	field, ok := doc.Fields["key"]
	if !ok {
		return
	}
	if field.Type != DocumentFieldTypeString {
		return
	}
	val, isString := field.Value.(string)
	if isString {
		documents[val] = doc
	}
}

// функція для отримання данних з мапи повертає два значення

func Get(key string) (*Document, bool) {
	doc, ok := documents[key]
	if !ok {
		return nil, false
	}
	return doc, ok
}

// функція для видалення данних з мапи

func Delete(key string) bool {
	_, ok := documents[key]
	if ok {
		delete(documents, key)
		return true
	}
	return false
}

// функція для отримання списку всіх данних у мапі

func List() []*Document {
	result := make([]*Document, 0, len(documents))
	for _, doc := range documents {
		result = append(result, doc)
	}
	return result
}
