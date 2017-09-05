<template>
  <div class="mdl-grid">
    <div class="mdl-cell mdl-cell--3-col"></div>
    <div class="mdl-cell mdl-cell--6-col">
      <div class="demo-list-action mdl-list">
        <div v-for="example in this.examples"
          class="mdl-list__item"
          @click="displayDetails(example.id)">
          <span class="mdl-list__item-primary-content">
            <i class="material-icons mdl-list__item-avatar">rowing</i>
                <span>{{ example.title }}</span>
          </span>
        </div>
      </div>
    </div>
    <div class="mdl-cell mdl-cell--3-col"></div>
    <router-link class="add-example-button mdl-button mdl-js-button mdl-button--fab mdl-button--colored" to="/examples/post">
      <i class="material-icons">add</i>
    </router-link>
  </div>
</template>

<script>
  import {HTTP} from '../../api'

  export default {
    methods: {
      displayDetails (id) {
        this.$router.push({name: 'Example Detail', params: { id: id }})
      }
    },
    data: () => ({
      examples: [],
      errors: []
    }),
    created () {
      HTTP.get('examples')
      .then(response => {
        this.examples = response.data
      })
      .catch(e => {
        this.errors.push(e)
      })
    }
  }
</script>

<style scoped>
  .add-example-button {
    position: fixed;
    right: 24px;
    bottom: 24px;
    z-index: 998;
  }

.mdl-list__item:hover {
  background-color: #eeeeee;
}


</style>
