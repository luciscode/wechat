package routers

import (
	"github.com/astaxie/beego"
	"web/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.Login{})
	beego.Router("/charts", &controllers.Charts{})

	beego.Router("/email1", &controllers.Email1{})

	beego.Router("/email2", &controllers.Email2{})

	beego.Router("/email3", &controllers.Email3{})

	beego.Router("/email4", &controllers.Email4{})

	beego.Router("/email5", &controllers.Email5{})

	beego.Router("/email6", &controllers.Email6{})

	beego.Router("/email5", &controllers.Email5{})

	beego.Router("/extra_404_option1", &controllers.Extra_404_option1{})
	beego.Router("/extra_404_option2", &controllers.Extra_404_option2{})
	beego.Router("/extra_404_option3", &controllers.Extra_404_option3{})
	beego.Router("/extra_500_option1", &controllers.Extra_500_option1{})
	beego.Router("/extra_500_option2", &controllers.Extra_500_option2{})
	beego.Router("/extra_faq", &controllers.Extra_faq{})
	beego.Router("/extra_image_manager", &controllers.Extra_image_manager{})
	beego.Router("/extra_invoice", &controllers.Extra_invoice{})

	beego.Router("/extra_lock", &controllers.Extra_lock{})
	beego.Router("/extra_pricing_table", &controllers.Extra_pricing_table{})
	beego.Router("/extra_profile", &controllers.Extra_profile{})
	beego.Router("/extra_search", &controllers.Extra_search{})
	beego.Router("/form_component", &controllers.Form_component{})
	beego.Router("/form_dropzone", &controllers.Form_dropzone{})
	beego.Router("/form_fileupload", &controllers.Form_fileupload{})
	beego.Router("/form_layout", &controllers.Form_layout{})
	beego.Router("/form_samples", &controllers.Form_samples{})
	beego.Router("/form_validation", &controllers.Form_validation{})
	beego.Router("/form_wizard", &controllers.Form_wizard{})
	beego.Router("/inbox", &controllers.Inbox{})
	beego.Router("/inbox_inbox", &controllers.Inbox_inbox{})
	beego.Router("/index", &controllers.Index{})
	beego.Router("/layout_ajax", &controllers.Layout_ajax{})
	beego.Router("/layout_ajax_content_2", &controllers.Layout_ajax_content_2{})
	beego.Router("/layout_ajax_content_3", &controllers.Layout_ajax_content_3{})
	beego.Router("/layout_blank_page", &controllers.Layout_blank_page{})
	beego.Router("/layout_boxed_not_responsive", &controllers.Layout_boxed_not_responsive{})
	beego.Router("/layout_boxed_page", &controllers.Layout_boxed_page{})
	beego.Router("/layout_email", &controllers.Layout_email{})
	beego.Router("/layout_horizontal_menu1", &controllers.Layout_horizontal_menu1{})
	beego.Router("/layout_horizontal_menu2", &controllers.Layout_horizontal_menu2{})
	beego.Router("/layout_horizontal_sidebar_menu", &controllers.Layout_horizontal_sidebar_menu{})
	beego.Router("/layout_promo", &controllers.Layout_promo{})
	beego.Router("/layout_sidebar_closed", &controllers.Layout_sidebar_closed{})
	beego.Router("/login", &controllers.Login{})
	beego.Router("/login_soft", &controllers.Login_soft{})
	beego.Router("/maps_google", &controllers.Maps_google{})
	beego.Router("/maps_vector", &controllers.Maps_vector{})
	beego.Router("/page_about", &controllers.Page_about{})
	beego.Router("/page_blog", &controllers.Page_blog{})
	beego.Router("/page_blog_item", &controllers.Page_blog_item{})
	beego.Router("/page_calendar", &controllers.Page_calendar{})
	beego.Router("/page_coming_soon", &controllers.Page_coming_soon{})
	beego.Router("/page_contact", &controllers.Page_contact{})
	beego.Router("/page_news", &controllers.Page_news{})
	beego.Router("/page_news_item", &controllers.Page_news_item{})
	beego.Router("/page_timeline", &controllers.Page_timeline{})
	beego.Router("/portlet_draggable", &controllers.Portlet_draggable{})
	beego.Router("/portlet_general", &controllers.Portlet_general{})
	beego.Router("/table_advanced", &controllers.Table_advanced{})
	beego.Router("/table_basic", &controllers.Table_basic{})
	beego.Router("/table_editable", &controllers.Table_editable{})
	beego.Router("/table_managed", &controllers.Table_managed{})
	beego.Router("/table_responsive", &controllers.Table_responsive{})
	beego.Router("/ui_buttons", &controllers.Ui_buttons{})
	beego.Router("/ui_general", &controllers.Ui_general{})
	beego.Router("/ui_jqueryui", &controllers.Ui_jqueryui{})
	beego.Router("/ui_modals", &controllers.Ui_modals{})
	beego.Router("/ui_nestable", &controllers.Ui_nestable{})
	beego.Router("/ui_sliders", &controllers.Ui_sliders{})
	beego.Router("/ui_tabs_accordions", &controllers.Ui_tabs_accordions{})
	beego.Router("/ui_tiles", &controllers.Ui_tiles{})
	beego.Router("/ui_tree", &controllers.Ui_tree{})
	beego.Router("/ui_typography", &controllers.Ui_typography{})

	//微信

	beego.Router("/all", &controllers.All{})
	beego.Router("/autoreply", &controllers.Autoreply{})
	beego.Router("/addautoreplytext", &controllers.Addautoreplytext{})
	beego.Router("/addautoreplynews", &controllers.Addautoreplynews{})
	beego.Router("/addautoreplynewses", &controllers.Addautoreplynewses{})
	beego.Router("/test", &controllers.Test{})
	beego.Router("/head", &controllers.Head{})

}
