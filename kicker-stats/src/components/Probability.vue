<template>
  <span>
    <v-banner id="topTitle">Calculate Probability</v-banner>
    <v-container fill-height fluid class="heightadjustment">
      <v-form class="FormStyling">
        <v-row>
          <v-col cols="12" sm="6" md="12"
            ><v-card-title>What are the odds?</v-card-title
            ><v-card-text
              >Have you ever wondered the odds of making a field goal at 60
              yards, well now you can find out. Provide the yardage and we can
              give you the odds of completion.</v-card-text
            ></v-col
          >
          <v-col cols="12" sm="6" md="12">
            <v-text-field label="Selected Kicker"></v-text-field
          ></v-col>
          <v-col cols="12" sm="6" md="12">
            <v-text-field label="What yardline"></v-text-field
          ></v-col>
          <v-col cols="12" sm="6" md="12">
            <!-- <v-card><v-input>Temp</v-input></v-card> -->
            <v-radio-group
              >What hash?<v-radio
                label="Right"
                color="white"
                value="right"
              ></v-radio>
              <v-radio label="Left" color="white" value="left"></v-radio>
            </v-radio-group>
          </v-col>
        </v-row>
        <v-btn class="btnSubmit">Tell Me the Odds</v-btn>
      </v-form>
    </v-container>
    <v-footer color="green darken-3">Contact info</v-footer>
  </span>
</template>

<script>
export default {
  name: "Probability",
};
</script>
<style scoped>
#topTitle {
  color: white;
  font-size: px;
  background-color: #4e342e;
  /* padding-top: 10px; */
  text-align: center;
}
form {
  background-color: #2e7d32;
  margin-left: 15%;
  margin-right: 15%;
  border-radius: 3%;
  padding: 10px;
  border-color: white;
  border-width: 3px;
}
/* .fieldImage {
  background-image: url("../assets/footballField.jpg");
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  max-height: 100vh;
} */
.btnSubmit {
  text-align: center;
  margin-top: 10px;
  margin-left: 40%;
  margin-right: 40%;
}
.heightadjustment {
  background-image: url("../assets/footballField.jpg");
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  max-height: 100vh;
}
</style>