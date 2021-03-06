// Code generated by reactGen. DO NOT EDIT.

package react

// TextAreaProps defines the properties for the <textarea> element
type TextAreaProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	ClassName               string
	Cols                    uint
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	DefaultValue string
	ID           string
	Key          string

	OnChange
	OnClick

	Placeholder string
	Ref
	Role  string
	Rows  uint
	Style *CSS
	Value string
}

func (t *TextAreaProps) assign(_v *_TextAreaProps) {

	_v.AriaExpanded = t.AriaExpanded

	_v.AriaHasPopup = t.AriaHasPopup

	_v.AriaLabelledBy = t.AriaLabelledBy

	_v.ClassName = t.ClassName

	_v.Cols = t.Cols

	_v.DangerouslySetInnerHTML = t.DangerouslySetInnerHTML

	if t.DataSet != nil {
		for dk, dv := range t.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if t.DefaultValue != "" {
		_v.DefaultValue = t.DefaultValue
	}

	if t.ID != "" {
		_v.ID = t.ID
	}

	if t.Key != "" {
		_v.Key = t.Key
	}

	if t.OnChange != nil {
		_v.o.Set("onChange", t.OnChange.OnChange)
	}

	if t.OnClick != nil {
		_v.o.Set("onClick", t.OnClick.OnClick)
	}

	_v.Placeholder = t.Placeholder

	if t.Ref != nil {
		_v.o.Set("ref", t.Ref.Ref)
	}

	_v.Role = t.Role

	_v.Rows = t.Rows

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = t.Style.hack()

	_v.Value = t.Value

}
