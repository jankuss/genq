package parser

func (t GenqTypeReference) IsCollectionType() bool {
  return t.Name == "List" || t.Name == "Set" || t.Name == "Map"
}
