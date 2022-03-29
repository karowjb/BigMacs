<template  class="background">
  <span>
    <v-banner id="topTitle">Top Ten Kickers</v-banner>
    <v-container fill-height fluid class="fieldImage">
      <v-form id="search">
        <v-row>
          <v-col cols="12" sm="6" md="12"
            ><v-card-title>Searching</v-card-title
            ><v-card-text
              >Please fill in the questions below to select the data you are
              looking for.</v-card-text
            ></v-col
          >
          <v-col cols="12" sm="6" md="12">
            <v-text-field label="Favorite Kicker"></v-text-field
          ></v-col>
          <v-col cols="12" sm="6" md="12">
            <v-text-field label="What season?"></v-text-field
          ></v-col>
          <v-col cols="12" sm="6" md="12">
            <v-card
              ><v-range-slider label="How many yards"
                >max="50" min="0"</v-range-slider
              ></v-card
            ></v-col
          >
          <v-col cols="12" sm="6" md="12">
            <v-card><v-input>Temp</v-input></v-card>
          </v-col>
        </v-row>
      </v-form>
      <v-container id="result">
        <v-simple-table>This is a table</v-simple-table>
        <br />
        <v-list>
          <v-list-item>One</v-list-item>
          <v-list-item>Two</v-list-item>
          <v-list-item>Three</v-list-item>
          <v-list-item>Four</v-list-item>
          <v-list-item>Five</v-list-item>
          <v-list-item>Six</v-list-item>
          <v-list-item>Seven</v-list-item>
          <v-list-item>Eight</v-list-item>
          <v-list-item>Nine</v-list-item>
          <v-list-item>Ten</v-list-item>
        </v-list>
      </v-container>
    </v-container>
    <v-footer color="green darken-3">Contact info</v-footer>
  </span>
</template>

<script>
export default {
  name: "TopTen",
};
</script>
<style scoped>
#search {
  width: 50%;
  height: 100%;
}
#result {
  width: 50%;
  height: 100%;
  padding: 5px;
  background-color: #2e7d32;
}
#topTitle {
  color: white;
  font-size: px;
  background-color: #4e342e;
  padding-top: 10px;
  text-align: center;
}
form {
  background-color: #2e7d32;
  /* margin-left: 15%;
  margin-right: 15%;
  border-radius: 3%; */
  padding: 5px;
  border-color: white;
  border-width: 3px;
}
.fieldImage {
  background-image: url("../assets/footballField.jpg");
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  max-height: 100vh;
}
</style>