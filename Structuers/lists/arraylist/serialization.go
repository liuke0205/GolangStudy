package arraylist

import "encoding/json"

// ToJSON 输出容器元素的JSON表示
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.elements[:list.size])
}

// MarshalJSON @implements json.Marshaler
func (list *List) MarshalJSON() ([]byte, error) {
	return list.ToJSON()
}

// FromJSON 从输入的JSON表示填充容器的元素
func (list *List) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &list.elements)
	if err == nil {
		list.size = len(list.elements)
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (list *List) UnmarshalJSON(data []byte) error {
	return list.FromJSON(data)
}
