<template>
  <v-app>
    <app-nav></app-nav>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>
<script>
import AppNav from '@/components/AppNav.vue';
import Vuetify from 'vuetify/lib';

const vuetify = new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#ecebf1',
        secondary: '#1f1e1f',
        accent: '#399c59',
        error: '#b71c1c',
        anchor: '#1f1e1f',
      },
    },
  },
})

export default ({
  name: "App",
  components: {
    AppNav,
  },
  vuetify,
})
</script>
<style scoped>

</style>