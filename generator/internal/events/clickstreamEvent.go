/*
 * MIT License
 *
 * Copyright (c) 2026 git-sudo-404
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * Of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * Copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * Copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package events

type ClickstreamEventType string

const (
	ClickPageView  ClickstreamEventType = "user.page_view"
	ClickClick     ClickstreamEventType = "user.click"
	ClickSearch    ClickstreamEventType = "user.search"
	ClickAddToCart ClickstreamEventType = "user.add_to_cart"
	ClickCheckout  ClickstreamEventType = "user.checkout_started"
	ClickPurchase  ClickstreamEventType = "user.purchase"
)

type ClickstreamEvent struct {
	EventMetaData

	SessionID string               `json:"session_id"`
	Action    ClickstreamEventType `json:"action"`
	Page      *string              `json:"page,omitempty"`
	Element   *string              `json:"element,omitempty"`
	ElementID *string              `json:"element_id,omitempty"`
	Metadata  map[string]string    `json:"metadata,omitempty"`
}

func (ce *ClickstreamEvent) SetSessionID(SessionID string) *ClickstreamEvent {
	ce.SessionID = SessionID
	return ce
}

func (ce *ClickstreamEvent) SetAction(Action ClickstreamEventType) *ClickstreamEvent {
	ce.Action = Action
	return ce
}

func (ce *ClickstreamEvent) SetPage(Page string) *ClickstreamEvent {
	ce.Page = &Page
	return ce
}

func (ce *ClickstreamEvent) SetElement(Element string) *ClickstreamEvent {
	ce.Element = &Element
	return ce
}

func (ce *ClickstreamEvent) SetElementID(ElementID string) *ClickstreamEvent {
	ce.ElementID = &ElementID
	return ce
}

func (ce *ClickstreamEvent) SetMetadata(Metadata map[string]string) *ClickstreamEvent {
	ce.Metadata = Metadata
	return ce
}
