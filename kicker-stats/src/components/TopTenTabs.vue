<template>
  <v-card>
    <v-tabs centered center-active dark>
      <v-tab v-for="tab in tabs" :key="tab.name">{{ tab.name }}</v-tab>
      <v-tab-item v-for="table in tables" :key="table.id">
        <v-card flat>
          <v-data-table
            :headers="headers"
            :items="table"
            hide-default-footer
          ></v-data-table>
        </v-card>
      </v-tab-item>
    </v-tabs>
  </v-card>
</template>
<script>
import axios from "axios";
const fgURL = "http://localhost:8000/toptenfg";
const lfgURL = "http://localhost:8000/toptenlongestfg";
const kieURL = "http://localhost:8000/toptenkintoendzone";

export default {
  name: "TopTenTabs",
  data() {
    return {
      tabs: [
        { name: "Field Goals Made" },
        { name: "Longest Field Goal" },
        { name: "Kickoffs Into Endzone" },
      ],
      tables: [],
      headers: [
        { text: "#", align: "center", sortable: false, value: "index" },
        { text: "Name", align: "start", sortable: false, value: "name" },
        {
          text: "Jersey Number",
          align: "center",
          sortable: false,
          value: "jerseyNum",
        },
        { text: "Team", sortable: false, value: "team" },
        {
          text: "Field Goals Made",
          align: "center",
          sortable: false,
          value: "fieldGoalsMade",
        },
        {
          text: "Longest Field Goal (yd)",
          align: "center",
          sortable: false,
          value: "longestFG",
        },
        {
          text: "Kickoffs Into Endzone",
          align: "center",
          sortable: false,
          value: "kickoffsEndzone",
        },
        { text: "Year", align: "center", sortable: false, value: "year" },
      ],
    };
  },
  async created() {
    try {
      //get table by field goals
      const response1 = await axios.post(fgURL);
      let respTable1 = response1.data.data;
      let fgtable = [];
      let index = 1;
      respTable1.forEach((entry) => {
        let newData = {
          index: index,
          name: entry.kickerfirstname + " " + entry.kickerlastname,
          jerseyNum: entry.jerseynum,
          team: entry.teamlocation + " " + entry.teamname,
          fieldGoalsMade: entry.fieldgoalsmade,
          longestFG: entry.longestfieldgoal,
          kickoffsEndzone: entry.endzonekickoffs,
          year: entry.seasonyear,
        };
        fgtable.push(newData);
        ++index;
      });
      this.tables.push(fgtable);

      //get table by longest fg
      const response2 = await axios.post(lfgURL);
      let respTable2 = response2.data.data;
      let lfgtable = [];
      index = 1;
      respTable2.forEach((entry) => {
        let newData = {
          index: index,
          name: entry.kickerfirstname + " " + entry.kickerlastname,
          jerseyNum: entry.jerseynum,
          team: entry.teamlocation + " " + entry.teamname,
          fieldGoalsMade: entry.fieldgoalsmade,
          longestFG: entry.longestfieldgoal,
          kickoffsEndzone: entry.endzonekickoffs,
          year: entry.seasonyear,
        };
        lfgtable.push(newData);
        ++index;
      });
      this.tables.push(lfgtable);

      //get tables by kickoffs into endzone
      const response3 = await axios.post(kieURL);
      let respTable3 = response3.data.data;
      let kietable = [];
      index = 1;
      respTable3.forEach((entry) => {
        let newData = {
          index: index,
          name: entry.kickerfirstname + " " + entry.kickerlastname,
          jerseyNum: entry.jerseynum,
          team: entry.teamlocation + " " + entry.teamname,
          fieldGoalsMade: entry.fieldgoalsmade,
          longestFG: entry.longestfieldgoal,
          kickoffsEndzone: entry.endzonekickoffs,
          year: entry.seasonyear,
        };
        kietable.push(newData);
        ++index;
      });
      this.tables.push(kietable);
    } catch (e) {
      console.log(e);
    }
  },
};
</script>
<style scoped>
v-card {
  background-color: var(--main-bg);
}
</style>
