<template>
    <span>
        <HomeHero></HomeHero>
        <div class="paragraph">
            <h2 class="text-h4 font-weight-bold text-center">A REVOLUTIONARY STATISTICS EXPERIENCE BUILT WITH REAL DATA</h2>
            <p class="text-p">Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Purus viverra accumsan in nisl nisi scelerisque eu. Sed ullamcorper morbi tincidunt ornare massa eget. Mauris commodo quis imperdiet massa tincidunt. Id velit ut tortor pretium. Egestas erat imperdiet sed euismod nisi. Proin libero nunc consequat interdum varius. Felis imperdiet proin fermentum leo vel orci. Viverra aliquet eget sit amet tellus cras adipiscing. Amet nisl purus in mollis. In egestas erat imperdiet sed euismod nisi porta lorem mollis.

            Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Purus viverra accumsan in nisl nisi scelerisque eu. Sed ullamcorper morbi tincidunt ornare massa eget. Mauris commodo quis imperdiet massa tincidunt. Id velit ut tortor pretium. Egestas erat imperdiet sed euismod nisi. Proin libero nunc consequat interdum varius. Felis imperdiet proin fermentum leo vel orci. Viverra aliquet eget sit amet tellus cras adipiscing. Amet nisl purus in mollis. In egestas erat imperdiet sed euismod nisi porta lorem mollis.
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
