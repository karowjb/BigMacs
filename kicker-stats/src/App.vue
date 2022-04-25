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
        primary: '#3f51b5',
        secondary: '#b0bec5',
        accent: '#8c9eff',
        error: '#b71c1c',
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