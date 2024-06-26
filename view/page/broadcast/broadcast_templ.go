// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package view_broadcast

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/wa-broadcast/view"
import "github.com/wa-broadcast/view/component"
import "github.com/wa-broadcast/view/partial/alert"

// script FormatPhone(phone, regionCode string) {

// }
func PhoneInput() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex gap-x-2\" x-data=\"countries\"><details x-ref=\"dropdown\" @click.outside=\"$refs.dropdown.open = false\" class=\"dropdown w-[25%]\"><summary tabindex=\"0\" class=\"btn flex gap-3 flex-nowrap\" @click=\"toggleContent\"><div :class=\"`fi fi-${selected.code.toLowerCase()}`\"></div><div x-text=\"&#39;+&#39; + selected.dialCodes[0]\"></div></summary><ul tabindex=\"0\" class=\"dropdown-content z-[1] menu shadow bg-base-100 rounded-box w-60 h-60 overflow-y-scroll flex-nowrap p-2\" x-data=\"{search: &#39;&#39;}\"><input type=\"text\" x-model=\"search\" placeholder=\"search dialcode\" class=\"input input-bordered w-full max-w-xs mb-2\"><template x-for=\"(country, i) in data\" :key=\"i\"><li @click=\"selected = country; $refs.dropdown.open = false\" x-show=\"!search || country.dialCodes[0] === search\" :class=\"{&#39;bg-slate-800/60 rounded-xl&#39;: selected.code === country.code}\"><div><div :class=\"`fi fi-${country.code.toLowerCase()}`\"></div><div x-text=\"&#39;+&#39; + country.dialCodes[0]\"></div></div></li></template></ul></details> <input @input=\"updatePhoneNumber($event, i-1)\" class=\"input input-bordered w-full\" placeholder=\"ex: 089234342341\"> <button x-show=\"i === field\" class=\"btn w-10\" type=\"button\" @click=\"field += 1\">+</button> <button x-show=\"i !== field\" class=\"btn w-10\" type=\"button\" @click=\"field -= 1\">-</button></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Broadcast() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = (view_component.Nav{make(map[string]string)}).Original().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <section class=\"container mx-auto flex min-h-screen flex-col justify-center items-center\"><div class=\"join join-vertical gap-5 w-1/2\" x-data=\"{ field: 1, phone: Array(this.field).join(&#39;.&#39;).split(&#39;.&#39;), msg: &#39;&#39;}\" hx-disinherit=\"*\"><div id=\"res\"></div><div id=\"phoneInput\" class=\"space-y-2\"><template x-for=\"i in field\" :key=\"i\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = PhoneInput().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</template></div><textarea x-model=\"msg\" class=\"textarea textarea-bordered join-item\" placeholder=\"Message\"></textarea> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, view_partial_alert.RemoveAlert(3000))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button hx-post=\"/wa/broadcast\" hx-trigger=\"click\" hx-swap=\"innerHTML transition:true\" hx-target=\"#res\" hx-on::after-request=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 templ.ComponentScript = view_partial_alert.RemoveAlert(3000)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var4.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-indicator=\"#loading-btn\" :hx-vals=\"JSON.stringify({&#39;numbers&#39;: phone, &#39;message&#39;: msg})\" class=\"btn\" type=\"button\">Send <span id=\"loading-btn\" class=\"loading loading-spinner loading-md htmx-indicator\"></span></button></div></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = view.Layout("broadcast-wa").Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
