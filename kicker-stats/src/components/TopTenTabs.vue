<template>
    <v-card>
        <v-tabs v-model="tab" centered dark>
            <v-tab v-for="tab in tabs" :key="tab.name">{{tab.name}}</v-tab>
            <v-tab-item v-for="table in tables" :key="table.id">
                <v-card flat>
                    <v-data-table :headers="headers" :items="desserts" hide-default-footer></v-data-table>
                </v-card>
            </v-tab-item>
        </v-tabs>
    </v-card>
</template>
<script>

//get tables from api

export default ({
    name: 'TopTenTabs',
    data() {
        return { 
            tabs: [
                {name: 'Field Goals Made'},
                {name: 'Single Yardage'},
                {name: 'Extra Points Made'}
            ],
            tables: [
                {table: 'table1', id: 0},
                {table: 'table2', id: 1},
                {table: 'table3', id: 2}
            ],
            headers: [
                { text: '#', sortable: false, value: 'index'},
                { text: 'Name', align: 'start', sortable: false, value: 'name' },
                { text: 'Jersey Number', sortable: false, value: 'jerseyNum' },
                { text: 'Team (Current)', sortable: false, value: 'team' },
                { text: 'Field Goals Made', sortable: false, value: 'fieldGoals' },
                { text: 'Single Yardage (yd)', sortable: false, value: 'yardage' },
                { text: 'Extra Points Made', sortable: false, value: 'extraPoints' },
            ]
        }
    }
})
</script>
<style scoped>
v-card {
    background-color: var(--main-bg);
}
</style>
