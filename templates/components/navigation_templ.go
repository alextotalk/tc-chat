// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Navigation() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"sticky-header\" class=\"fixed top-0 z-[60] flex items-center justify-center w-full h-16 duration-500 ease-out bg-white border-b bg-opacity-90 backdrop-blur-md border-neutral-400 border-opacity-20\"><div class=\"flex items-center justify-between w-full px-4 mx-auto 2xl:px-0 max-w-7xl\"><div class=\"relative z-10 flex items-center w-auto leading-10 lg:flex-grow-0 lg:flex-shrink-0 lg:text-left\"><button @click=\"open = ! open\" class=\"flex items-center justify-center w-8 h-8 mr-3 border rounded-md lg:hidden border-neutral-200 text-neutral-800 hover:text-black hover:bg-neutral-100 active:bg-white\"><svg class=\"w-5 h-5\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><g fill=\"none\" stroke=\"none\"><path d=\"M2.74902 6.75C2.74902 5.09315 4.09217 3.75 5.74902 3.75H18.2507C19.9075 3.75 21.2507 5.09315 21.2507 6.75V17.25C21.2507 18.9069 19.9075 20.25 18.2507 20.25H5.74902C4.09217 20.25 2.74902 18.9069 2.74902 17.25V6.75Z\" stroke=\"currentColor\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path><path d=\"M13.75 3.75V20.25\" stroke=\"currentColor\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path><path d=\"M16.75 7.75L18.25 7.75\" stroke=\"currentColor\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path><path d=\"M16.75 11L18.25 11\" stroke=\"currentColor\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path><path d=\"M16.75 14.25L18.25 14.25\" stroke=\"currentColor\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></g></svg></button> <a href=\"/pines\" class=\"inline-flex sm:mr-8 items-end font-sans text-2xl font-extrabold text-left text-black no-underline bg-transparent cursor-pointer group focus:no-underline\"><svg class=\"w-auto lg:h-6 h-5\" viewBox=\"0 0 355 99\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"><defs><path d=\"M119.1 87V66.4h19.8c34.3 0 34.2-49.5 0-49.5-11 0-22 .1-33 .1v70h13.2zm19.8-32.7h-19.8V29.5h19.8c16.8 0 16.9 24.8 0 24.8zm32.6-30.5c0 9.5 14.4 9.5 14.4 0s-14.4-9.5-14.4 0zM184.8 87V37.5h-12.2V87h12.2zm22.8 0V61.8c0-7.5 5.1-13.8 12.6-13.8 7.8 0 11.9 5.7 11.9 13.2V87h12.2V61.1c0-15.5-9.3-24.2-20.9-24.2-6.2 0-11.2 2.5-16.2 7.4l-.8-6.7h-10.9V87h12.1zm72.1 1.3c7.5 0 16-2.6 21.2-8l-7.8-7.7c-2.8 2.9-8.7 4.6-13.2 4.6-8.6 0-13.9-4.4-14.7-10.5h38.5c1.9-20.3-8.4-30.5-24.9-30.5-16 0-26.2 10.8-26.2 25.8 0 15.8 10.1 26.3 27.1 26.3zM292 56.6h-26.6c1.8-6.4 7.2-9.6 13.8-9.6 7 0 12 3.2 12.8 9.6zm41.2 32.1c14.1 0 21.2-7.5 21.2-16.2 0-13.1-11.8-15.2-21.1-15.8-6.3-.4-9.2-2.2-9.2-5.4 0-3.1 3.2-4.9 9-4.9 4.7 0 8.7 1.1 12.2 4.4l6.8-8c-5.7-5-11.5-6.5-19.2-6.5-9 0-20.8 4-20.8 15.4 0 11.2 11.1 14.6 20.4 15.3 7 .4 9.8 1.8 9.8 5.2 0 3.6-4.3 6-8.9 5.9-5.5-.1-13.5-3-17-6.9l-6 8.7c7.2 7.5 15 8.8 22.8 8.8z\" id=\"a\"></path></defs><g fill=\"none\" fill-rule=\"evenodd\"><g fill=\"currentColor\"><path d=\"M19.742 49h28.516L68 83H0l19.742-34z\"></path><path d=\"M26 69h14v30H26V69zM4 50L33.127 0 63 50H4z\"></path></g><g fill-rule=\"nonzero\"><use fill=\"currentColor\" xlink:href=\"#a\"></use><use fill=\"currentColor\" xlink:href=\"#a\"></use></g></g></svg> <span class=\"flex opacity-90 group-hover:scale-150 group-hover:opacity-100 items-center h-full group-hover:-rotate-6 ease-out duration-500 px-0.5 py-px ml-2 -translate-x-px text-[0.65rem] font-bold leading-none border-[2px] rounded border-black -translate-y-px\">UI</span></a><nav class=\"items-center hidden space-x-5 text-sm font-medium lg:flex\"><a class=\"relative transition-colors text-neutral-700 hover:text-black group\" href=\"/pines/docs/introduction\"><span>Documentation</span> <span class=\"absolute bottom-0 w-0 h-0.5 duration-300 ease-out bg-black group-hover:w-full left-1/2 group-hover:left-0 group-hover:-translate-x-0\"></span></a> <a class=\"relative transition-colors text-neutral-700 hover:text-black group\" href=\"/pines/docs/accordion\"><span>Elements</span> <span class=\"absolute bottom-0 w-0 h-0.5 duration-300 ease-out bg-black group-hover:w-full left-1/2 group-hover:left-0 group-hover:-translate-x-0\"></span></a> <a onclick=\"if (!window.__cfRLUnblockHandlers) return false; window.dispatchEvent(new CustomEvent(&#39;open-playground&#39;, { detail: {}}));\" class=\"relative transition-colors text-neutral-700 hover:text-black group\" href=\"javascript:;\"><span>Playground</span> <span class=\"absolute bottom-0 w-0 h-0.5 duration-300 ease-out bg-black group-hover:w-full left-1/2 group-hover:left-0 group-hover:-translate-x-0\"></span></a> <a class=\"relative hidden transition-colors text-neutral-700 hover:text-black group\" href=\"/pines/examples\"><span>Examples</span> <span class=\"absolute bottom-0 w-0 h-0.5 duration-300 ease-out bg-black group-hover:w-full left-1/2 group-hover:left-0 group-hover:-translate-x-0\"></span></a> <a class=\"relative transition-colors text-neutral-700 hover:text-black group\" target=\"_blank\" href=\"https://github.com/thedevdojo/pines\"><span>Github</span> <span class=\"absolute bottom-0 w-0 h-0.5 duration-300 ease-out bg-black group-hover:w-full left-1/2 group-hover:left-0 group-hover:-translate-x-0\"></span></a></nav></div><div class=\"relative space-x-1 font-medium leading-10 sm:space-x-2 md:flex-grow-0 md:flex-shrink-0 md:text-right lg:flex-grow-0 lg:flex-shrink-0\"><button @click=\"fullscreenModal = ! fullscreenModal\" class=\"inline-flex items-center px-3 sm:px-5 text-sm border-0 rounded-md cursor-pointer h-9 focus:outline-none md:mt-0 text-neutral-900 hover:text-neutral-800 hover:bg-neutral-100\">Login</button> <button onclick=\"if (!window.__cfRLUnblockHandlers) return false; document.getElementById(&#39;auth&#39;).dispatchEvent(new CustomEvent(&#39;open&#39;, {}));\" target=\"_blank\" class=\"inline-flex items-center px-3 sm:px-5 text-sm border-0 rounded-md cursor-pointer h-9 focus:outline-none md:mt-0 bg-neutral-900 hover:bg-neutral-800 hover:text-white text-gray-100\">Signup</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = NavigationTab().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LogIn().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
