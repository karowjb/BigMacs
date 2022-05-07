<template>
    <v-card>
        <v-tabs centered dark>
            <v-tab v-for="tab in tabs" :key="tab.name">{{tab.name}}</v-tab>
            <v-tab-item v-for="table in tables" :key="table.id">
                <v-card flat>
                    <v-data-table :headers="headers" :items="table" hide-default-footer></v-data-table>
                </v-card>
            </v-tab-item>
        </v-tabs>
    </v-card>
</template>
<script>
import axios from 'axios';
const fgURL = 'http://localhost:8000/toptenfg';
const lfgURL = 'http://localhost:8000/toptenyardage';
const epmURL = 'http://localhost:8000/toptenepm';

export default ({
    name: 'TopTenTabs',
    data() {
        return { 
            tabs: [
                {name: 'Field Goals Made'},
                {name: 'Longest Field Goal'},
                {name: 'Extra Points Made'}
            ],
            tables: [],
            headers: [
                { text: '#', align: 'center', sortable: false, value: 'index'},
                { text: 'Name', align: 'start', sortable: false, value: 'name' },
                { text: 'Jersey Number', align: 'center', sortable: false, value: 'jerseyNum' },
                { text: 'Team (Current)', sortable: false, value: 'team' },
                { text: 'Field Goals Made', align: 'center', sortable: false, value: 'fieldGoals' },
                { text: 'Longest Field Goal (yd)', align: 'center', sortable: false, value: 'longest' },
                { text: 'Extra Points Made', sortable: false, value: 'extraPoints' },
                { text: 'Year', align: 'center', sortable: false, value: 'year' },
            ]
        }
    },
    async created() {
        try {
            //get table by field goals
            const response1 = await axios.post(fgURL);
            let respTable1 = response1.data.data;
            let fgtable = [];
            console.log(respTable1);
            let index = 1;
            respTable1.forEach(entry => {
                let newData = {
                    'index': index,
                    'name': entry.kickerfirstname + ' ' + entry.kickerlastname,
                    'jerseyNum': entry.jerseynum,
                    'team': entry.teamlocation + ' ' + entry.teamname,
                    'fieldGoals': entry.fieldgoalsmade,
                    'longest': entry.longestfieldgoal,
                    'extraPoints': 'Loading...',
                    'year': entry.seasonyear
                }
                fgtable.push(newData);
                ++index;
            });
            console.log(fgtable);
            this.tables.push(fgtable);

            //get table by single yardage
            const response2 = await axios.post(lfgURL);
            let respTable2 = response2.data.data;
            let lfgtable = [];
            console.log(respTable2);
            index = 1;
            respTable2.forEach(entry => {
                let newData = {
                    'index': index,
                    'name': entry.kickerfirstname + ' ' + entry.kickerlastname,
                    'jerseyNum': entry.jerseynum,
                    'team': entry.teamlocation + ' ' + entry.teamname,
                    'fieldGoals': entry.fieldgoalsmade,
                    'longest': 'Loading...',
                    'extraPoints': 'Loading...',
                    'year': entry.seasonyear
                }
                lfgtable.push(newData);
                ++index;
            });
            this.tables.push(lfgtable);

            //get tables by estra points made
            const response3 = await axios.post(epmURL);
            let epmtable = response3.data.data;
            this.tables.push(epmtable);

        } catch(e) {
            console.log(e);
        }
    }
})
</script>
<style scoped>
v-card {
    background-color: var(--main-bg);
}
</style>
