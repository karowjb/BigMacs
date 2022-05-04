<template>
<span>
  <PageHeader header="- TOP 10 -"></PageHeader>
  <div class='main'>
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
.main {
  background-image: linear-gradient(180deg, rgba(1,35,37,1) 0%, rgba(236,235,241,1) 100%);
}
</style>