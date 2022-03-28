<template>
    <v-container class='home-hero' fluid fill-height>
    <v-row justify='center' align-content='center'>
        <v-col>
            <div class="text-h1 white--text font-weight-bold text-center">Welcome</div>
            <div class="text-h3 white--text text-center">- TO THE -</div>
            <div class="text-h1 white--text font-weight-bold text-center">Game Changer</div>
        </v-col>
    </v-row>
</v-container>
</template>

<script>
export default ({
    name: "HomeHero"
})
</script>

<style scoped>
.home-hero {
    background-image: url("../assets/home-hero.jpg");
    background-size: cover;
    background-repeat: no-repeat;
    background-position: center;
    text-shadow: 6px 6px 4px #002d2d;
}
</style>