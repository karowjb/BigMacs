<template>
  <v-card>
    <v-card-text>
      <v-form>
        <v-text-field v-model="playerName" label="Name">
        </v-text-field>
        <v-select v-model="team"
          :items="teams"
          label="Team"
        ></v-select>
        <v-range-slider
          v-model="heightRange"
          :min="48"
          :max="100"
          :color="blue"
          :track-color="primary"
          label="Height"
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
                @change="$set(range, 1, $event)"
              ></v-text-field>
            </template>
        </v-range-slider>
        <v-slider
          v-model="weight"
          :min="150"
          :max="300"
          label="Weight (lb)"
          class="align-center"
          ><template v-slot:append>
            <v-text-field
              v-model="weight"
              class="mt-0 pt-0"
              type="number"
              style="width: 60px"
            ></v-text-field>
          </template>
        </v-slider>
        <v-btn>Search</v-btn>
        <v-btn>Reset</v-btn>
      </v-form>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "SearchInput",
  data() {
    return {
      teams: [],
      team: '',
      playerName: '',
      heightRange: [48, 100],
    }
  },
  methods: {
    
  }
  //get teams from api
};
</script>
<style scoped>

</style>