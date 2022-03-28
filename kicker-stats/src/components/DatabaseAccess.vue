<template  class="background">
  <div class="fieldImage">
    <v-banner id="topTitle">Search Database</v-banner>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" sm="6" md="12">
            <v-card-title>Searching</v-card-title>
            <v-card-text>Please fill in the questions below to select the data you arelooking for.</v-card-text>
            </v-col>
          <v-col cols="12" sm="6" md="12">
            <v-text-field label="Favorite Kicker"></v-text-field>
            </v-col>
          <v-col cols="12" sm="6" md="12">
            <v-text-field label="What season?"></v-text-field>
            </v-col>
          <v-col cols="12" sm="6" md="12"> </v-col>
          <v-col cols="12" sm="6" md="12">
            <v-card><v-input>Temp</v-input></v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
    <v-footer color="green darken-3">Contact info</v-footer>
  </div>
</template>

<script>
export default {
  name: "Searching",
};
</script>
<style scoped>
#topTitle {
  color: white;
  font-size: px;
  background-color: #4e342e;
  padding-top: 10px;
  text-align: center;
}
form {
  background-color: #2e7d32;
  margin: 50px;
  border-radius: 3%;
  border-color: white;
  border-width: 3px;
}
.fieldImage {
  background-image: url("../assets/footballField.jpg");
}
</style>