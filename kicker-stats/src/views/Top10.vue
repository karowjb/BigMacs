<template>
<span>
  <PageHeader header="- Top 10 -"></PageHeader>
  <div class='mainbg'>
    <v-container>
      <TopTenTabs></TopTenTabs>
    </v-container>
  </div>
</span>
</template>
<script>
import PageHeader from "@/components/PageHeader.vue"
import TopTenTabs from "@/components/TopTenTabs.vue"
export default {
  name: "Top10",
  components: { PageHeader, TopTenTabs }
};
</script>
<style scoped>

</style>