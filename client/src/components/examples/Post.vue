<template>
  <form>
    <div class="mdl-grid">
      <div class="mdl-cell mdl-cell--4-col mdl-cell--8-col-tablet">
        <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label is-upgraded is-dirty">
          <input id="title" v-model="title" type="text" class="mdl-textfield__input"/>
          <label for="title" class="mdl-textfield__label">Title</label>
        </div>
        <div class="mdl-textfield mdl-js-textfield">
          <textarea class="mdl-textfield__input" type="text" rows= "3" v-model="body" id="body" ></textarea>
          <label class="mdl-textfield__label" for="body">Your Example Body...</label>
        </div>
        <ul v-if="errors && errors.length">
          <li v-for="error of errors">
            <span>{{error.message}}</span>
          </li>
        </ul>
        <div class="actions">
          <a @click.prevent="postExample()" class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored">
            POST EXAMPLE
          </a>
        </div>
      </div>
    </div>
  </form>
</template>
<script>
  import {HTTP} from '../../api'

  export default {
    methods: {
      postExample () {
        HTTP.post('examples', {
          title: this.title,
          body: this.body
        })
        .then(response => {
          this.$router.push({name: 'Example Detail', params: { id: response.data.id }})
        })
        .catch(e => {
          this.errors.push(e)
        })
      }
    },
    data: () => ({
      title: '',
      body: '',
      errors: []
    })

  }
</script>
<style scoped>
  .waiting {
    padding: 10px;
    color: #555;
  }
</style>
