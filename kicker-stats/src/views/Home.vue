<template>
    <span>
        <HomeHero></HomeHero>
        <div class="paragraph">
            <h2 class="text-h4 font-weight-bold text-center">A REVOLUTIONARY STATISTICS EXPERIENCE BUILT WITH REAL DATA</h2>
            <p class="text-p">
                The Game Changer is the home for anything NFL Kicker statistics related. Whether you are wondering where the soccer ball is in this version of “football” or are watching NFL Redzone for 10 hours every Sunday during the season, there is a page on this site for you. Need to know if the kicker on your significant other’s favorite team is any good? Just look him up! Take a gander at the Top 10 page to see how this kicker stacks up to others. Need to find a kicker for your fantasy football team? Look no further than the Top 10 page. Anxious about a last second kick during a timeout? Head on over to the probabilities page to see whether your anxiety is justified. What are you waiting for? Go look up some stats on your favorite kicker today!
            </p>
        </div>
    </span>
</template>
<script>
import HomeHero from "@/components/HomeHero.vue";

export default ({
    name: 'Home',
    components: { HomeHero }
})
</script>
<style scoped>
.paragraph {
    margin: 30px 60px;
}
.paragraph h2 {
    margin-bottom: 10px;
    color: #555555;
    max-height: 100vh;
}
</style>
