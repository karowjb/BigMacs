<template>
  <v-app>
    <AppNav></AppNav>
    <v-main>
      <router-view></router-view>
    </v-main>
    <AppFooter></AppFooter>
  </v-app>
</template>
<script>
import Vue from "vue";
import AppNav from "@/components/AppNav.vue";
import AppFooter from "@/components/AppFooter.vue";
import Vuetify from "vuetify/lib";
import axios from "axios";

const vuetify = new Vuetify({
  theme: {
    themes: {
      light: {
        primary: "#ecebf1",
        secondary: "#1f1e1f",
        accent: "#399c59",
        darkgreen: "#003a38",
        error: "#b71c1c",
        anchor: "#1f1e1f",
      },
      dark: {
        primary: "#1f1e1f",
        secondary: "#399c59",
        accent: "#ecebf1",
      },
    },
  },
});
Vue.mixin({
  methods: {
    inToFt: function (num) {
      return Math.floor(num / 12) + "'" + (num % 12) + '"';
    },
  },
});

export default {
  name: "App",
  components: {
    AppNav,
    AppFooter,
  },
  vuetify,
  axios,
};
</script>
<style>
:root {
  --main-bg: #ecebf1;
  --main-text: #1f1e1f;
  --main-accent: #399c59;
}
.mainbg {
  background-image: linear-gradient(
    180deg,
    rgba(1, 35, 37, 1) 0%,
    rgba(255, 255, 255, 1) 100%
  );
}
</style>