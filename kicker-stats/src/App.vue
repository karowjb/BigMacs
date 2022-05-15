<template>
  <v-app>
    <AppNav></AppNav>
    <v-main>
      <router-view></router-view>
    </v-main>
    <!-- <AppFooter></AppFooter> -->
  </v-app>
</template>
<script>
import AppNav from "@/components/AppNav.vue";
// import AppFooter from "@/components/AppFooter.vue";
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

export default {
  name: "App",
  components: {
    AppNav,
  },
  vuetify,
  axios,
  created() {
    document.title = "BigMacs Sports Stats";
  }
};
</script>
<style>
.mainbg {
  background-image: linear-gradient(180deg, rgba(1, 35, 37, 1) 0%, rgba(255, 255, 255, 1) 100%);
}
.homebg {
  background-image: linear-gradient(180deg, rgba(57, 76, 88, 1) 0%, rgba(255, 255, 255, 1) 100%);
}
.extraTopBotMarg {
    margin-top: 32px;
    margin-bottom: 32px;
}
.extraBotPad {
  padding-bottom: 64px;
}
</style>