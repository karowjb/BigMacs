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
            <v-card-title class="justify-center">{{this.probFG}}%</v-card-title>
        </v-card-text>
    </v-card>
</template>
<script>
import axios from 'axios';
const getKickersURL = 'http://localhost:8000/getkickers';
const probabilitiesURL = 'http://localhost:8000/getprobabilities';
// const probKurl = '';

export default ({
    name: 'ProbabilityInput',
    data() {
        return {
            kicker: {},
            kickerNames: [],
            kickers: [],
            probFG: 0,
            probK: 0,
        }
    },
    methods: {
        searchProb() {
            try {
                let id = '';
                if (this.kickerNames.indexOf(this.kicker) != 0) {
                    let x = this.kickerNames.indexOf(this.kicker) - 1;
                    id = this.kickers[x].kickerid;
                    console.log(id);
                } else {
                    return;
                }
                axios.post(probabilitiesURL, id).then((response) => {
                    this.kicker = response.data.data;

                    this.probFG = this.kicker.fieldgoalsmade/this.kicker.fieldgoalattempts;
                    this.probK = this.kickerkickoffyards/this.kicker.kickoffs;
                    console.log(this.probFG + '%');
                    console.log(this.probK);
                },
                (error) => {
                    console.log(error);
                });
            } catch (e) {
             console.log(e);
            }
        }
    },
    async created() {
        try {
        //get kickers
        const response = await axios.post(getKickersURL);
        this.kickers = response.data.data;
        this.kickers.forEach(kicker => {
            this.kickerNames.push(kicker.kickerfirstname + ' ' + kicker.kickerlastname);
        });
        } catch(e) {
        console.log(e);
        }
    },
})
</script>
<style scoped>
.form {
    text-align: center;
}
</style>
