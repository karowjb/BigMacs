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
                <br><br>
                <v-row>
                    <v-col cols='12' sm='6'>
                        <v-card-title class='justify-center'>Chance of Making a Field Goal:</v-card-title>
                        <v-card-title class='justify-center black--text results animeFG' v-show='validFG'>{{this.probFG}}%</v-card-title>
                        <v-card-title class='justify-center red--text' v-show='!validFG'>Not Enough Available Data</v-card-title>
                    </v-col>
                    <v-col cols='12' sm='6'>
                        <v-card-title class='justify-center'>Average Kickoff Yardage:</v-card-title>
                        <v-card-title class='justify-center black--text results animeK' v-show='validK'>{{this.probK}} yd.</v-card-title>
                        <v-card-title class='justify-center red--text' v-show='!validK'>Not Enough Available Data</v-card-title>
                    </v-col>
                </v-row>
        </v-card-text>
    </v-card>
</template>
<script>
import axios from 'axios';
const getKickersURL = 'http://localhost:8000/getkickers';
const probabilitiesURL = '/probabilities';

export default ({
    name: 'ProbabilityInput',
    data() {
        return {
            kicker: '',
            kickerNames: [],
            kickers: [],
            probFG: 0,
            probK: 0,
            validFG: true,
            validK: true,
            animeFG: null,
            animeK: null,
        }
    },
    methods: {
        searchProb() {
            try {
                let id = this.kickers[this.kickerNames.indexOf(this.kicker)].kickerid;                
                let k = {
                    'kickerid' : id
                }
                axios.post(probabilitiesURL, k).then(
                    (response) => {
                    let resp = response.data.data;
                    let kicker = {
                        totalk: resp[0].totalattemptsk,
                        totalkyards: resp[0].totalyardsk,
                        totalfg: resp[0].totalattemptsfg,
                        totalfgmade: resp[0].totalmadefg
                    }
                    let probFG = ((kicker.totalfgmade/kicker.totalfg)*100);
                    let probK = kicker.totalkyards/kicker.totalk;
                    this.probFG = probFG;
                    this.probK = probK;

                    if (kicker.totalfg === 0) {
                        this.validFG = false;
                    } else {
                        this.validFG = true;
                        this.animeFG = this.$anime({
                            targets: '.animeFG',
                            duration: 3000,
                            easing: 'cubicBezier(.4,.34,.1,1)',
                            innerText: this.probFG + '%',
                            round: 100,
                        })
                    }
                    if (kicker.totalk === 0) {
                        this.validK = false;
                    } else {
                        this.validK = true;
                        this.animeK = this.$anime({
                            targets: '.animeK',
                            duration: 3000,
                            easing: 'cubicBezier(.4,.34,.1,1)',
                            innerText: this.probK + ' yd.',
                            round: 100,
                        })
                    }
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
.results {
    font-size: 2em;
}
</style>
