<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm6 offset-sm3>
        <v-progress-linear :indeterminate="true" v-show="loading"></v-progress-linear>
        <v-alert v-show="examples.length < 1 && !loading" color="info" icon="info" value="true">
        No Examples yet...   (click that little green add button)
        </v-alert>
        <v-card v-show="errors.length < 1">
          <v-list three-line>
            <v-subheader>Examples</v-subheader>
            <template v-for="(example, index) in examples">
              <v-divider :key="index"></v-divider>
              <v-list-tile v-bind:key="index" @click="displayDetails(example.id)">
                <v-list-tile-content>
                  <v-list-tile-title v-html="example.title"></v-list-tile-title>
                  <v-list-tile-sub-title v-html="example.body"></v-list-tile-sub-title>
                </v-list-tile-content>
              </v-list-tile>
            </template>
          </v-list>
        </v-card>
        <v-alert color="error" v-show="errors.length > 0" icon="warning" value="true">
          <p v-for="(error, i) in errors" :key="i">
            {{ error.message }}
          </p>
        </v-alert>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { HTTP } from '../../api'

export default {
  methods: {
    displayDetails(id) {
      this.$router.push({ name: 'Example Detail', params: { id: id } })
    }
  },
  data: () => ({
    examples: [],
    errors: [],
    loading: false
  }),
  created() {
    this.loading = true
    HTTP.get('examples')
      .then(response => {
        this.examples = response.data
        this.loading = false
      })
      .catch(e => {
        this.errors.push(e)
        this.loading = false
      })
  }
}
</script>
