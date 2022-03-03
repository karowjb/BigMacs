<template>
<span>
    <v-navigation-drawer app v-model='drawer'>
        <v-list>
            <template v-for='(link, index) in pages'>
                <v-list-item :key='index' :to='link.url'>
                    <v-list-item-content>{{link.title}}</v-list-item-content>
                </v-list-item>
                <v-divider :key='`div-${index}`'></v-divider>
            </template>
        </v-list>
    </v-navigation-drawer>
    <v-app-bar color='light-blue accent-4' dense app>
        <v-app-bar-nav-icon class='hidden-md-and-up' @click='drawer = !drawer'></v-app-bar-nav-icon>
        <v-img src='@/assets/donut.png' max-width=40></v-img>
        <router-link to='/'>
        <v-app-bar-title>BigMacs Sports Stats</v-app-bar-title>
        </router-link>
        <v-spacer></v-spacer>
        <v-btn to='/' class='hidden-sm-and-down' plain>Home</v-btn>
        <v-btn to='/top10' class='hidden-sm-and-down' plain>Top 10</v-btn>
        <v-btn to='/search' class='hidden-sm-and-down' plain>Search</v-btn>
    </v-app-bar>
</span>
</template>
<script>
export default ({
    name: 'Nav',
    data() {
        return {
            drawer: false,
            pages: [
                {link: "Home", url: "/"},
                {link: "Top 10", url: "/top10"},
                {link: "Search", url: "/search"},
                {link: "Probabilities", url: "/probabilities"}
            ]
        }
    }
})
</script>
<style scoped>
a, a:visited {
    color: black;
    text-decoration: none;
    margin-left: 10px;
}
</style>
