<template>
    <v-card>
        <v-card-text class='form'>
            <v-form>
                <v-select v-model="kicker"
                    :items="kickerNames"
                    label="Kicker"
                    selected-color="green darken-1"
                    required
                ></v-select>
                <v-btn v-on:click="searchProb" color="green darken-1">Calculate</v-btn>
                </v-form>
            <v-card-title class="justify-center">100%</v-card-title>
        </v-card-text>
    </v-card>
</template>
<script>
export default ({
    name: 'ProbabilityInput',
    data() {
        return {
            kicker: '',
            kickerNames: [],
            probFG: 0,
            probK: 0
        }
    },
    methods: {
        searchProb() {
            let id = '';
            if (this.kickerNames.indexOf(this.kicker) != 0) {
                let x = this.kickerNames.indexOf(this.kicker) - 1;
                id = this.kickers[x].kickerid;
            }
            id;
        }
    },
})
</script>
<style scoped>
.form {
    text-align: center;
}
</style>
