<template>
<v-card style="backgroundColor: red" id='expandedCardBG' flat>
    <v-card class='expandedCard' flat>
        <v-card-text>
            <v-row>
                <v-col cols='4'>Height:<br>{{player.height}}</v-col>
                <v-col cols='4'>Weight:<br>{{player.weight}} lb.</v-col>
                <v-col cols='4'>Home:<br>{{player.teamHomeStadium}}</v-col>
            </v-row>
            <v-row>
                <v-col cols='4'>Kickoffs:<br>{{player.kickoffs}}</v-col>
                <v-col cols='4'>Fieldgoal Attempts:<br>{{player.fieldgoalAttempts}}</v-col>
                <v-col cols='4'>Fieldgoals Made:<br>{{player.fieldgoalsMade}}</v-col>
            </v-row>
            <v-row>
                <v-col cols='4'>Total Fieldgoal Yards:<br>{{player.fieldgoalYards}} yd</v-col>
                <v-col cols='4'>Longest Fieldgoal:<br>{{player.fieldgoalLongest}} yd</v-col>
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
