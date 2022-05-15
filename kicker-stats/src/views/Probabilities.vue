<template>
<span>
  <PageHeader header="- Probabilities -"></PageHeader>
  <div class='mainbg extraBotPad'>
    <v-container>
      <ProbabilityInput></ProbabilityInput>
    </v-container>
  </div>
</span>
</template>
<script>
import PageHeader from "@/components/PageHeader.vue";
import ProbabilityInput from "@/components/ProbabilityInput.vue";
export default {
  name: "Probabilities",
  components: {
    ProbabilityInput,
    PageHeader,
  },
};
</script>