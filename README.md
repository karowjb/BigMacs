# NFL Kicker Statistics

![home_page](sample_pictures/home_page.png)

The goal of this project was to create a website using Vue/Vuetify with database or API calls integrated as well. Specifically, we chose to make a website that interacted with a database of NFL kicker statistics. The website is split into three main pages: Top 10, Search, and Probabilities.

## The Database

The database is static with data from the 2021-2022 NFL season. We explored using an API to gather our data, but this was unsuccessful because the best free API would not work when running locally. This led us to use the API to populate our own database and use Go to make calls to the database. 

There are three main tables in the database: Teams, Kickers, and Kicker seasons. These three tables all have keys that reference other tables. The main connections are relating a kicker to a team and a kicker season to a kicker. The 