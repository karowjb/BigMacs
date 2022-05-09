<template>
    <v-card>
    <v-card-text>
        <v-data-table 
            :headers="searchHeaders"
            :items="players"
            :items-per-page="10"
            :single-expand="singleExpand"
            :expanded.sync="expanded"
            item-key="id"
            show-expand>
            <template v-slot:expanded-item="{ headers, item }">
                <td :colspan="headers.length">More info about {{ item.name }}
                </td>
            </template>
        </v-data-table>
    </v-card-text>
  </v-card>
</template>
<script>

//get results from api

export default ({
    name: "SearchOutput",
    data() {
        return {
            searchHeaders: [
                { text: 'Name', align: 'start', value: 'name' },
                { text: 'Team', value: 'team'},
                { text: 'Jersey Number', value: 'jerseyNum'}
            ],
            players: [],
        }
    },
    props: ['searchResults'],
    watch: {
        searchResults: function() {
            console.log("Changed");
            let results = [];
            this.searchResults.forEach(v => {
                let height = inToFt(v.kickerheight);
                let newData = {
                    'id': v.kickerid,
                    'name': v.kickerfirstname + ' ' + v.kickerlastname,
                    'height': height,
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
            console.log(results)
            console.log(this.players);
        },
    },
})
</script>
