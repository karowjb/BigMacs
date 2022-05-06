<template>
<span>
  <PageHeader header="- Search Database -"></PageHeader>
  <div class='mainbg'>
    <v-container>
      <v-row>
        <v-col cols="6">
          <SearchInput></SearchInput>
        </v-col>
        <v-col>
          <SearchOutput></SearchOutput>
        </v-col>
      </v-row>
    </v-container>
  </div>
</span>
</template>
<script>
import PageHeader from "@/components/PageHeader.vue";
import SearchInput from "@/components/SearchInput.vue";
import SearchOutput from "@/components/SearchOutput.vue";

export default {
  name: "Search",
  components: {
    PageHeader,
    SearchInput,
    SearchOutput
  },
};
</script>