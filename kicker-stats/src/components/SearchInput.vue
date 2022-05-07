<template>
  <v-card>
    <v-card-text>
      <v-form>
        <v-text-field v-model="fname" label="First Name"></v-text-field>
        <v-text-field v-model="lname" label="Last Name"></v-text-field>
        <v-select v-model="team"
          :items="teamNames"
          label="Team"
          selected-color="green darken-1"
        ></v-select>
        <v-range-slider
          v-model="heightRange"
          :min="48"
          :max="100"
          color="green darken-1"
          track-color="primary"
          label="Height (in)"
          class="align-center"
          ><template v-slot:prepend>
              <v-text-field
                :value="heightRange[0]"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 60px"
                @change="$set(heightRange, 0, $event)"
              ></v-text-field>
            </template>
            <template v-slot:append>
              <v-text-field
                :value="heightRange[1]"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 60px"
                @change="$set(heightRange, 1, $event)"
              ></v-text-field>
            </template>
        </v-range-slider>
        <v-range-slider
          v-model="weightRange"
          :min="150"
          :max="300"
          color="green darken-1"
          track-color="primary"
          label="Weight (lb)"
          class="align-center"
          ><template v-slot:prepend>
              <v-text-field
                :value="weightRange[0]"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 60px"
                @change="$set(weightRange, 0, $event)"
              ></v-text-field>
            </template>
            <template v-slot:append>
              <v-text-field
                :value="weightRange[1]"
                class="mt-0 pt-0"
                hide-details
                single-line
                type="number"
                style="width: 60px"
                @change="$set(weightRange, 1, $event)"
              ></v-text-field>
            </template>
        </v-range-slider>
        <v-btn v-on:click="searchdb" color="green darken-1" id="searchbtn">Search</v-btn>
        <v-btn v-on:click="resetInputs" color="amber">Reset</v-btn>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script>
import axios from 'axios';
const getTeamsURL = 'http://localhost:8000/getteams';

export default {
  name: "SearchInput",
  data() {
    return {
      teamNames: [''],
      teams: [],

      //inputs
      team: '',
      fname: '',
      lname: '',
      heightRange: [48, 100],
      weightRange: [150, 300],
    }
  },
  methods: {
    searchdb() {
      let newSearch = {
        'kickerteamid': this.team,
        'kickerfirstname': this.fname,
        'kickerlastname': this.lname,
        'kickerheightmin': this.heightRange[0],
        'kickerheightmax': this.heightRange[1],
        'kickerweightmin': this.weightRange[0],
        'kickerweightmax': this.weightRange[1],
      }
      if (this.fname.trim().length() != 0) {
        newSearch.kickerfirstname = this.fname;
      }
      if (this.lname.trim().length() != 0) {
        newSearch.kickerlastname = this.lname;
      }
      if (this.team != this.teamNames[0]) {
        let index = this.teamNames.indexOf(this.team) - 1;
        newSearch.kickerteamid = this.teams[index].Team_id;
      }


    },
    resetInputs() {
      this.team = this.teamNames[0];
      this.fname = '';
      this.lname = '';
      this.heightRange = [48, 100];
      this.weightRange = [150, 300];
    },
  },
  async created() {
    try {
      //get teams
      const response = await axios.post(getTeamsURL);
      this.teams = response.data.data;
      this.teams.forEach(team => {
        this.teamNames.push(team.teamlocation + ' ' + team.teamname);
      });

    } catch(e) {
      console.log(e);
    }
  }
}
</script>
<style scoped>
#searchbtn {
  margin-right: 16px;
}
</style>