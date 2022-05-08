<template>
    <v-card>
    <v-card-text>
        <v-data-table 
            :headers="searchHeaders"
            :items="players"
            :items-per-page="20"
            :single-expand="singleExpand"
            :expanded.sync="expanded"
            item-key="name"
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
    // computed: {
    //     showResults() {
    //         this.players = this.searchResults;
    //     }
    // },
    watch: {
        searchResults: function() {
            console.log("Changed");
            this.players = this.searchResults;
            console.log(this.searchResults)
            console.log(this.players);
        },
    },
})
</script>
