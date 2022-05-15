<template>
    <span>
        <HomeHero></HomeHero>
        <div class='homebg'>
            <v-container>
                <HomeDescription></HomeDescription>
            </v-container>
        </div>
    </span>
</template>
<script>
import HomeHero from "@/components/HomeHero.vue";
import HomeDescription from "@/components/HomeDescription.vue";

export default ({
    name: 'Home',
    components: { HomeHero, HomeDescription }
})
</script>
<style scoped>

</style>
