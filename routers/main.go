package routers

import (
  "strconv"
  "github.com/kataras/iris"
  "github.com/alsmile/goApiGateway/utils"
  "github.com/alsmile/goApiGateway/controllers"
)

func Start() {
  app := iris.New()

  app.OnErrorCode(iris.StatusNotFound, controllers.NotFound)

  admin := app.Party(utils.GlobalConfig.Domain.AdminDomain)
  {
    admin.StaticWeb("/assets", "./admin/web/dist/assets")

    admin.Get("/", controllers.Index)
    admin.Get("/browser.html", controllers.Browser)
    admin.Get("/captcha", controllers.Captcha)

    admin.Post("/api/login", controllers.Login)
    admin.Post("/api/signup", controllers.SignUp)
    admin.Post("/api/sign/active", controllers.SignActive)
    admin.Post("/api/forget/password", controllers.ForgetPassword)
    admin.Post("/api/sign/new/password", controllers.NewPassword)
    admin.Get("/api/user/profile", controllers.UserProfile)

    admin.Get("/api/sign/config", controllers.GetSignConfig)

    admin.Get("/api/site/list", controllers.Auth, controllers.SiteList)
    admin.Get("/api/site/get", controllers.Auth, controllers.SiteGet)
    admin.Post("/api/site/save", controllers.Auth, controllers.SiteSave)
    admin.Post("/api/site/api/save", controllers.Auth, controllers.SiteApiSave)
    admin.Get("/api/site/api/get", controllers.Auth, controllers.SiteApiGet)
    admin.Get("/api/site/api/del", controllers.Auth, controllers.SiteApiDel)
    admin.Get("/api/site/api/list", controllers.Auth, controllers.SiteApiList)

    admin.Post("/api/site/api/list/by/domains", controllers.Auth, controllers.SiteApiListByDomains)

    admin.Any("/api/test", controllers.ProxyTest)
  }

  app.Any("/{url:path}", controllers.ProxyDo)

  strPort := strconv.Itoa(int(utils.GlobalConfig.Domain.Port))
  app.Run(iris.Addr(":" + strPort))
}