<template>
  <v-container fluid>
    <v-layout row>
      <v-flex>
        <div class="hello">
                <img src="../assets/logo.png" alt="Chive">
          <h1>{{ msg }}</h1>
          <p>Chive is a boilerplate for full-stack web development using Go, Chi, MongoDB, and Vue. When this project is built for deployment, everything is wrapped up into a single binary file.</p>
  
          <v-btn color="primary" 
            href="https://github.com/tutley/chive" 
            target="_blank" 
            small round>
            <v-icon left>code</v-icon>
            Github
          </v-btn>
          <p>&nbsp;</p>
          <h2>Ecosystem</h2>
          <v-btn color="secondary"
              href="https://github.com/go-chi/chi" target="_blank"
              small flat>Chi</v-btn>
            <v-btn color="secondary"
              href="https://golang.org" target="_blank"
              small flat>Go</v-btn>
            <v-btn color="secondary"
              href="https://www.mongodb.com" target="_blank"
              small flat>MongoDB</v-btn>
            <v-btn color="secondary"
              href="https://vuejs.org/" target="_blank"
              small flat>Vue</v-btn>
            <v-btn color="secondary"
              href="https://vuetifyjs.com" target="_blank"
              small flat>Vuetify</v-btn>
         
        </div>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: 'hello',
  data() {
    return {
      msg: 'Welcome to Chive'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #35495e;
}

.hello {
  text-align: center;
}
</style>
