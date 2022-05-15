import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify'
import VueAnime from 'vue-animejs';
 
Vue.use(VueAnime)

Vue.config.productionTip = false
Vue.use(vuetify, {
  theme: {
    themes: {
      light: {
        primary: '#ecebf1',
        secondary: '#1f1e1f',
        accent: '#399c59',
        error: '#b71c1c',
        anchor: '#1f1e1f',
      },
      dark: {
        primary: '#1f1e1f',
        secondary: '#399c59',
        accent: '#ecebf1',
      },
    },
  },
})

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
