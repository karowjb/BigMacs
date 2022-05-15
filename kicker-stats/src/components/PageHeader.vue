<template>
<span>
    <v-container class='container' fluid>
    <v-row justify='center' align-text='center'>
        <v-col>
            <div class="header text-h2 font-weight-bold hidden-xs-only">{{header}}</div>
            <div class="headerSm text-h4 font-weight-bold hidden-sm-and-up">{{header}}</div>
        </v-col>
    </v-row>
    </v-container>
</span>
</template>
<script>
export default ({
    name: 'PageHeader',
    props: ['header'],
})
</script>
<style scoped>
.container {
    text-align: center;
    background-image: url('../assets/page-header.jpg');
    background-size: cover;
    background-position: center;
}
.header {
    color: white;
    text-shadow: 6px 6px 4px #002d2d;
}
.headerSm {
    color: white;
    text-shadow: 3px 3px 2px #002d2d;
}



</style>
