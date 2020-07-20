module.exports = {
  /*
   ** Headers of the page
   */
  head: {
    title: "console.infinimesh.app",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: "Console Infinimesh UI"
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  plugins: ["@/plugins/ant-design-vue", "@/plugins/typeface-exo"],
  /*
   ** Customize the progress bar color
   */
  loading: { color: "#3B8070" },
  /*
   ** Build configuration
   */
  build: {
    /*
     ** Run ESLint on save
     */
    extend(config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: "pre",
          test: /\.(js|vue)$/,
          loader: "eslint-loader",
          exclude: /(node_modules)/
        });
      }
    }
  },
  buildModules: ["@nuxt/typescript-build"],
  modules: ["@nuxtjs/axios", "@nuxtjs/auth"],
  axios: {
    baseURL: "https://api.infinimesh.app/"
  },
  auth: {
    strategies: {
      local: {
        endpoints: {
          login: {
            url: "account/token",
            method: "post",
            propertyName: "token"
          },
          user: { url: "account", method: "get", propertyName: false },
          logout: false
        },
        tokenType: "bearer"
      }
    }
  },
  router: {
    middleware: ["auth"]
  },
  build: {
    loaders: {
      less: {
        javascriptEnabled: true,
        modifyVars: require("./assets/styles/antThemeMod.js")
      }
    }
  },
  css: [{ lang: "less", src: "@/assets/styles/themes.less" }]
};