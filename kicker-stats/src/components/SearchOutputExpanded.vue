<template>
<v-card id='expandedCardBG' flat>
    <v-card class='expandedCard' flat>
        <v-card-text>
            <v-row>
                <v-col cols='4'>Height: {{player.height}}</v-col>
                <v-col cols='4'>Weight: {{player.weight}} lb.</v-col>
                <v-col cols='4'>Home: {{player.teamHomeStadium}}</v-col>
            </v-row>
            <v-row>
                <v-col cols='4'>Kickoffs: {{player.kickoffs}}</v-col>
                <v-col cols='4'>Fieldgoal Attempts: {{player.fieldgoalAttempts}}</v-col>
            </v-row>
        </v-card-text>
    </v-card>
</v-card>
</template>
<script>
export default ({
    name: 'SearchOutputExpanded',
    data() {
        return {

        }
    },
    props: {
        player: {},
    },
    mounted() {
        let bg = document.getElementById('expandedCardBG');
        bg.style.background = 'linear-gradient(180deg, '+this.player.teamPrimaryColor+' 0%, '+this.player.teamSecondaryColor+' 100%)';
    }
})
</script>
<style scoped>
#expandedCardBG {
    margin: 0;
}
.expandedCard {
    margin-top: 16px;
    margin-bottom: 16px;
}
</style>
