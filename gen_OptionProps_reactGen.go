// Code generated by reactGen. DO NOT EDIT.

package react

// OptionProps defines the properties for the <option> element
type OptionProps struct {
	AriaExpanded            bool
	AriaHasPopup            bool
	AriaLabelledBy          string
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	DataSet
	ID  string
	Key string

	OnChange
	OnClick

	Ref
	Role  string
	Style *CSS
	Value string
}

func (o *OptionProps) assign(_v *_OptionProps) {

	_v.AriaExpanded = o.AriaExpanded

	_v.AriaHasPopup = o.AriaHasPopup

	_v.AriaLabelledBy = o.AriaLabelledBy

	_v.ClassName = o.ClassName

	_v.DangerouslySetInnerHTML = o.DangerouslySetInnerHTML

	if o.DataSet != nil {
		for dk, dv := range o.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if o.ID != "" {
		_v.ID = o.ID
	}

	if o.Key != "" {
		_v.Key = o.Key
	}

	if o.OnChange != nil {
		_v.o.Set("onChange", o.OnChange.OnChange)
	}

	if o.OnClick != nil {
		_v.o.Set("onClick", o.OnClick.OnClick)
	}

	if o.Ref != nil {
		_v.o.Set("ref", o.Ref.Ref)
	}

	_v.Role = o.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = o.Style.hack()

	_v.Value = o.Value

}
