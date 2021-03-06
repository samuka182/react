// Code generated by reactGen. DO NOT EDIT.

package react

// CodeProps defines the properties for the <code> element
type CodeProps struct {
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
}

func (c *CodeProps) assign(_v *_CodeProps) {

	_v.AriaExpanded = c.AriaExpanded

	_v.AriaHasPopup = c.AriaHasPopup

	_v.AriaLabelledBy = c.AriaLabelledBy

	_v.ClassName = c.ClassName

	_v.DangerouslySetInnerHTML = c.DangerouslySetInnerHTML

	if c.DataSet != nil {
		for dk, dv := range c.DataSet {
			_v.o.Set("data-"+dk, dv)
		}
	}

	if c.ID != "" {
		_v.ID = c.ID
	}

	if c.Key != "" {
		_v.Key = c.Key
	}

	if c.OnChange != nil {
		_v.o.Set("onChange", c.OnChange.OnChange)
	}

	if c.OnClick != nil {
		_v.o.Set("onClick", c.OnClick.OnClick)
	}

	if c.Ref != nil {
		_v.o.Set("ref", c.Ref.Ref)
	}

	_v.Role = c.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	_v.Style = c.Style.hack()

}
