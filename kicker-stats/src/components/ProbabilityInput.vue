<template>
    <v-card>
        <v-card-text>
            <v-form>
                <v-text-field v-model="kickerName" label="Name"></v-text-field>
            </v-form>
            <v-card-title class="justify-center">100%</v-card-title>
        </v-card-text>
    </v-card>
</template>
<script>
export default ({
    name: 'ProbabilityInput',
    data() {
        return {
            kickerName: '',
        }
    }
})
</script>
