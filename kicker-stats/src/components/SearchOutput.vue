<template>
    <v-card>
    <v-card-text>
        <v-data-table 
            :headers="searchHeaders"
            :items="players"
            :items-per-page="10"
            :single-expand=false
            :expanded.sync="expanded"
            item-key="id"
            show-expand>
            <template v-slot:expanded-item="{ headers, item }">
                <td :colspan="headers.length" :style="{background: 'linear-gradient(180deg, '+item.teamPrimaryColor+' 0%, '+item.teamSecondaryColor+' 100%)'}">
                    <SearchOutputExpanded :player=item></SearchOutputExpanded>
                </td>
            </template>
        </v-data-table>
    </v-card-text>
    </v-card>
</template>
<script>
import SearchOutputExpanded from '@/components/SearchOutputExpanded.vue';

export default ({
    name: "SearchOutput",
    components: { SearchOutputExpanded },
    data() {
        return {
            searchHeaders: [
                { text: 'Name', align: 'start', value: 'name' },
                { text: 'Team', value: 'teamName'},
                { text: 'Jersey Number', value: 'jerseyNum'}
            ],
            players: [],
            expanded: [],
        }
    },
    props: ['searchResults'],
    watch: {
        searchResults: function() {
            let results = [];
            if (Object.keys(this.searchResults).length === 0) {
                this.players = [];
                this.expanded = [];
                return;
            }
            this.searchResults.forEach(v => {
                let height = v.kickerheight;
                let formatHeight = Math.floor(height/12)+"'"+height%12+'"';
                let newData = {
                    'id': v.kickerid,
                    'name': v.kickerfirstname + ' ' + v.kickerlastname,
                    'jerseyNum': v.jerseynum,
                    'height': formatHeight,
                    'weight': v.kickerweight,
                    'teamId': v.kickerteamid,
                    'teamName': v.kickerteamlocation + ' ' + v.kickerteamname,
                    'teamHomeStadium': v.kickerteamhomestadium,
                    'teamPrimaryColor': v.kickerteamprimary,
                    'teamSecondaryColor': v.kickerteamsecondary,
                    'kickoffs': v.kickoffs,
                    'fieldgoalAttempts': v.fieldgoalsattempts,
                    'fieldgoalsMade': v.fieldgoalsmade,
                    'fieldgoalYards': v.fieldgoalyards,
                    'fieldgoalLongest': v.fieldgoallongest,
                    'xpAttempts': v.xpattempts,
                    'xpMade': v.xpmade
                }
                results.push(newData);
            })
            this.players = results;
        },
    },
})
</script>
