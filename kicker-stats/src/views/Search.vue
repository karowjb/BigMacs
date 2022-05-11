<template>
<span>
  <PageHeader header="- Search Database -"></PageHeader>
  <div class='mainbg'>
    <v-container>
      <v-row>
        <v-col cols="6">
          <SearchInput @search="search" @resetOutput="resetOutput"></SearchInput>
        </v-col>
        <v-col>
          <SearchOutput :searchResults=this.newSearch></SearchOutput>
        </v-col>
      </v-row>
    </v-container>
  </div>
</span>
</template>
<script>
import axios from 'axios';
import PageHeader from "@/components/PageHeader.vue";
import SearchInput from "@/components/SearchInput.vue";
import SearchOutput from "@/components/SearchOutput.vue";

const searchURL = '/searchplayers';

export default {
  name: "Search",
  components: {
    PageHeader,
    SearchInput,
    SearchOutput
  },
  data() {
    return {
      newSearch: {},
    }
  },
  methods: {
    search(search) {
        //api request to search with parameters
        try {
          axios.post(searchURL, search).then((response) => {
            let outputTable = response.data.data;
            // console.log(outputTable);
            this.newSearch = outputTable;
          }, (error) => {
            console.log(error);
          });
        } catch(e) {
          console.log(e);
        }
      
    },
    resetOutput(r) {
      if (r == true) {
        this.newSearch = {};
        console.log(this.newSearch);

      }
    }
  }
};
</script>