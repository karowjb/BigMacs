<template>
<span>
    <v-container class='homeHero hidden-xs-only' fluid fill-height>
        <v-row justify='center' align-content='center'>
            <v-col>
                <div class="text-h1 white--text font-weight-bold text-center">Welcome</div>
                <div class="text-h3 white--text text-center">- TO THE -</div>
                <div class="text-h1 white--text font-weight-bold text-center">Game Changer</div>
            </v-col>
        </v-row>
    </v-container>
    <v-container class='homeHero homeHeroSm hidden-sm-and-up'>
        <v-row justify='center' align-content='center'>
            <v-col>
                <div class="text-h3 white--text font-weight-bold text-center">Welcome</div>
                <div class="text-h4 white--text text-center">- TO THE -</div>
                <div class="text-h3 white--text font-weight-bold text-center">Game Changer</div>
            </v-col>
        </v-row>
    </v-container>
</span>
</template>

<script>
export default ({
    name: "HomeHero"
})
</script>

<style scoped>
.homeHero {
    background-image: url("../assets/home-hero.jpg");
    background-size: cover;
    background-repeat: no-repeat;
    background-position: center;
    text-shadow: 6px 6px 4px #002d2d;
}
.homeHeroSm {
    text-shadow: 3px 3px 2px #002d2d;
}
</style>