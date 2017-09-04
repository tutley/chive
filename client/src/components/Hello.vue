<template>
  <div class="hello">
          <img src="../assets/logo.png" alt="Chive">
    <h1>{{ msg }}</h1>
    <p>Chive is a boilerplate for full-stack web development using Go, Chi, MongoDB, and Vue. When this project is built for deployment, everything is wrapped up into a single binary file.</p>
    <p><a href="https://github.com/tutley/chive" target="_blank" rel="noopener">Github</a></p>
    <h2>Ecosystem</h2>
    <ul>
      <li><a href="https://github.com/go-chi/chi" target="_blank" rel="noopener">Chi</a></li>
      <li><a href="https://golang.org" target="_blank" rel="noopener">Go</a></li>
      <li><a href="https://www.mongodb.com" target="_blank" rel="noopener">MongoDB</a></li>
      <li><a href="http://vuejs.org/" target="_blank" rel="noopener">Vue</a></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'hello',
  data () {
    return {
      msg: 'Welcome to Chive'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
h1, h2 {
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
  color: #35495E;
}

.hello {
  text-align: center;
}
</style>
